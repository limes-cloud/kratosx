package web

import (
	"bytes"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	thttp "github.com/go-kratos/kratos/v2/transport/http"
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
}

func WithPreloadExts(exts map[string]bool) Option {
	return func(o *options) {
		o.preloadExts = exts
	}
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

func (c *cache) preload(dir string, exts map[string]bool) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		if exts[strings.ToLower(filepath.Ext(path))] {
			data, err := os.ReadFile(path)
			if err == nil {
				c.Set(path, data)
			}
		}
		return nil
	})
}

func Server(dir string, srv *thttp.Server, opts ...Option) {
	o := &options{preloadExts: defaultPreloadExts}
	for _, opt := range opts {
		opt(o)
	}

	c := newCache()
	c.preload(dir, o.preloadExts)

	content, ok := c.Get(filepath.Join(dir, "index.html"))
	if !ok {
		panic("index.html not found")
	}

	absDir, _ := filepath.Abs(dir)

	srv.HandleFunc("/", handler(dir, absDir, content, c))
	srv.HandleFunc("/{path:.*.*}", handler(dir, absDir, content, c))
}

func handler(dir, absDir string, content []byte, c *cache) http.HandlerFunc {
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
