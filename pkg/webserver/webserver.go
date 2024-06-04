package webserver

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"unsafe"

	"github.com/go-kratos/kratos/v2/log"
)

func Run(dir string, addr string, data map[string]any) {
	tpl := template.New("index")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(dir, r.URL.Path)
		if stat, err := os.Stat(path); err == nil && !stat.IsDir() {
			http.ServeFile(w, r, path)
			return
		}

		accept := r.Header.Get("Accept")
		if strings.Contains(accept, "text/html") {
			indexFilePath := filepath.Join(dir, "index.html")
			content, err := os.ReadFile(indexFilePath)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			parser, err := tpl.Parse(*(*string)(unsafe.Pointer(&content)))
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusOK)

			if err = parser.Execute(w, data); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
		}
		http.NotFound(w, r)
	})

	log.Infof("Starting web server at %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Warnf("Failed to start web server: %v", err)
	}
}
