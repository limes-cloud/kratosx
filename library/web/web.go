package web

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	thttp "github.com/go-kratos/kratos/v2/transport/http"
)

func Server(dir string, srv *thttp.Server) {
	path := filepath.Join(dir, "index.html")
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	srv.HandleFunc("/", handler(dir, content))
	srv.HandleFunc("/{path:.*.*}", handler(dir, content))
}

func handler(dir string, content []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tempIndex := content
		path := filepath.Join(dir, r.URL.Path)
		if stat, err := os.Stat(path); err == nil && !stat.IsDir() {
			http.ServeFile(w, r, path)
			return
		}

		accept := r.Header.Get("Accept")
		if strings.Contains(accept, "text/html") {
			if r.URL.Path != "/" {
				tempPath := path
				if stat, err := os.Stat(path); err == nil && stat.IsDir() {
					tempPath = filepath.Join(tempPath, "index.html")
				}

				tic, err := os.ReadFile(tempPath)
				if err == nil {
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
