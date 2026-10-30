package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	thttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx/config"
)

var defaultPreloadExts = map[string]bool{
	".html":  true,
	".css":   true,
	".js":    true,
	".json":  true,
	".svg":   true,
	".ico":   true,
	".png":   true,
	".jpg":   true,
	".jpeg":  true,
	".gif":   true,
	".woff":  true,
	".woff2": true,
	".ttf":   true,
	".eot":   true,
	".map":   true,
}

type Option func(*options)

type options struct {
	preloadExts map[string]bool
	injections  []*config.WebInjection
}

func WithPreloadExts(exts map[string]bool) Option {
	return func(o *options) {
		o.preloadExts = exts
	}
}

// WithInjections 在静态 HTML 加载到内存时批量替换标记。
func WithInjections(injections []*config.WebInjection) Option {
	return func(o *options) {
		o.injections = injections
	}
}

type injection struct {
	marker  string
	snippet string
}

var scriptMarkerPattern = regexp.MustCompile(`^<!--\s*([A-Za-z_$][A-Za-z0-9_$]*)\s*-->$`)

func prepareInjections(items []*config.WebInjection) ([]injection, error) {
	result := make([]injection, 0, len(items))
	for idx, item := range items {
		if item == nil {
			return nil, fmt.Errorf("injection[%d] is nil", idx)
		}

		marker := strings.TrimSpace(item.Marker)
		if marker == "" {
			return nil, fmt.Errorf("injection[%d].marker is required", idx)
		}

		var snippet string
		switch strings.ToLower(strings.TrimSpace(item.Type)) {
		case "script":
			matches := scriptMarkerPattern.FindStringSubmatch(marker)
			if len(matches) != 2 {
				return nil, fmt.Errorf("injection[%d].marker %q must contain a JavaScript identifier, for example <!--__APP_CONFIG__-->", idx, marker)
			}
			payload, err := json.Marshal(item.Value)
			if err != nil {
				return nil, fmt.Errorf("injection[%d].value: %w", idx, err)
			}
			snippet = fmt.Sprintf("<script>window.%s=%s;</script>", matches[1], payload)
		case "html":
			value, ok := item.Value.(string)
			if !ok {
				return nil, fmt.Errorf("injection[%d].value must be a string when type is html", idx)
			}
			snippet = value
		case "":
			return nil, fmt.Errorf("injection[%d].type is required (script or html)", idx)
		default:
			return nil, fmt.Errorf("injection[%d].type %q is not supported (script or html)", idx, item.Type)
		}

		result = append(result, injection{marker: marker, snippet: snippet})
	}
	return result, nil
}

func applyInjections(content []byte, injections []injection) ([]byte, map[string]bool) {
	found := make(map[string]bool, len(injections))
	result := content
	for _, item := range injections {
		if !bytes.Contains(result, []byte(item.marker)) {
			continue
		}
		found[item.marker] = true
		result = bytes.ReplaceAll(result, []byte(item.marker), []byte(item.snippet+item.marker))
	}
	return result, found
}

type cache struct {
	mu   sync.RWMutex
	data map[string][]byte
}

func newCache() *cache {
	return &cache{data: make(map[string][]byte)}
}

func (c *cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.data[key]
	return v, ok
}

func (c *cache) Set(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = val
}

func (c *cache) preload(dir string, exts map[string]bool, injections []injection) map[string]bool {
	found := make(map[string]bool, len(injections))
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		if exts[ext] {
			data, err := os.ReadFile(path)
			if err == nil {
				if ext == ".html" {
					injected, hits := applyInjections(data, injections)
					for marker := range hits {
						found[marker] = true
					}
					data = injected
				}
				c.Set(path, data)
			}
		}
		return nil
	})
	return found
}

func Server(dir string, srv *thttp.Server, opts ...Option) {
	o := &options{preloadExts: defaultPreloadExts}
	for _, opt := range opts {
		opt(o)
	}
	injections, err := prepareInjections(o.injections)
	if err != nil {
		panic(fmt.Sprintf("web injection: %v", err))
	}

	c := newCache()
	found := c.preload(dir, o.preloadExts, injections)

	content, ok := c.Get(filepath.Join(dir, "index.html"))
	if !ok {
		panic("index.html not found")
	}
	for _, item := range injections {
		if !found[item.marker] {
			panic(fmt.Sprintf("web injection marker %q not found in %s", item.marker, dir))
		}
	}

	absDir, _ := filepath.Abs(dir)

	srv.HandleFunc("/", handler(dir, absDir, content, c, injections))
	srv.HandleFunc("/{path:.*.*}", handler(dir, absDir, content, c, injections))
}

func handler(dir, absDir string, content []byte, c *cache, injections []injection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tempIndex := content
		path := filepath.Join(dir, r.URL.Path)
		absPath, _ := filepath.Abs(path)
		if !strings.HasPrefix(absPath, absDir) {
			http.NotFound(w, r)
			return
		}
		if stat, err := os.Stat(path); err == nil && !stat.IsDir() {
			data, ok := c.Get(path)
			if !ok {
				data, err = os.ReadFile(path)
				if err != nil {
					http.Error(w, "internal server error", http.StatusInternalServerError)
					return
				}
				if strings.EqualFold(filepath.Ext(path), ".html") {
					data, _ = applyInjections(data, injections)
				}
				c.Set(path, data)
			}
			http.ServeContent(w, r, filepath.Base(path), stat.ModTime(), bytes.NewReader(data))
			return
		}

		accept := r.Header.Get("Accept")
		if strings.Contains(accept, "text/html") {
			if r.URL.Path != "/" {
				tempPath := path
				if stat, err := os.Stat(path); err == nil && stat.IsDir() {
					tempPath = filepath.Join(tempPath, "index.html")
				}

				tic, ok := c.Get(tempPath)
				if ok {
					tempIndex = tic
				} else if data, err := os.ReadFile(tempPath); err == nil {
					tempIndex, _ = applyInjections(data, injections)
					c.Set(tempPath, tempIndex)
				}
			}

			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(tempIndex)
			return
		}

		http.NotFound(w, r)
	}
}
