package webserver

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"unsafe"

	"github.com/go-kratos/kratos/v2/log"
)

func Run(dir string, addr string, data map[string]any) {
	path := filepath.Join(dir, "index.html")
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if len(data) != 0 {
		var out = bytes.NewBuffer([]byte(""))
		tpl := template.New("html")
		parser, err := tpl.Parse(*(*string)(unsafe.Pointer(&content)))
		if err != nil {
			panic(err)
		}

		if err = parser.Execute(out, data); err != nil {
			panic(err)
		}
		content = out.Bytes()
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(dir, r.URL.Path)
		if stat, err := os.Stat(path); err == nil && !stat.IsDir() {
			http.ServeFile(w, r, path)
			return
		}

		accept := r.Header.Get("Accept")
		if strings.Contains(accept, "text/html") {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(content)
			return
		}
		http.NotFound(w, r)
	})

	log.Infof("Starting web server at %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic("Failed to start web server: " + err.Error())
	}
}
