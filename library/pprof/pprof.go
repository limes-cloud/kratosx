package pprof

import (
	"net/http"
	"net/http/pprof"

	thttp "github.com/go-kratos/kratos/v2/transport/http"

	"github.com/limes-cloud/kratosx/config"
)

func Server(conf *config.Pprof, srv *thttp.Server) {
	secret := conf.Secret
	query := conf.Query

	srv.HandleFunc("/debug/pprof", pprofServer(pprof.Index, secret, query))
	srv.HandleFunc("/debug/cmdline", pprofServer(pprof.Cmdline, secret, query))
	srv.HandleFunc("/debug/profile", pprofServer(pprof.Profile, secret, query))
	srv.HandleFunc("/debug/symbol", pprofServer(pprof.Symbol, secret, query))
	srv.HandleFunc("/debug/symbol", pprofServer(pprof.Symbol, secret, query))
	srv.HandleFunc("/debug/trace", pprofServer(pprof.Trace, secret, query))
	srv.HandleFunc("/debug/allocs", pprofServer(pprof.Handler("allocs").ServeHTTP, secret, query))
	srv.HandleFunc("/debug/block", pprofServer(pprof.Handler("block").ServeHTTP, secret, query))
	srv.HandleFunc("/debug/goroutine", pprofServer(pprof.Handler("goroutine").ServeHTTP, secret, query))
	srv.HandleFunc("/debug/heap", pprofServer(pprof.Handler("heap").ServeHTTP, secret, query))
	srv.HandleFunc("/debug/mutex", pprofServer(pprof.Handler("mutex").ServeHTTP, secret, query))
	srv.HandleFunc("/debug/threadcreate", pprofServer(pprof.Handler("threadcreate").ServeHTTP, secret, query))
}

func pprofServer(handler http.HandlerFunc, secret, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Query().Get(query) == secret {
			handler.ServeHTTP(writer, request)
		}
	}
}
