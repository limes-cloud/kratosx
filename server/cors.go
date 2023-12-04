package server

import (
	thttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	ec "github.com/limes-cloud/kratosx/config"
	"time"
)

func Cors(conf *ec.Cors) thttp.ServerOption {
	maxAge := time.Minute * 10
	if conf.MaxAge != 0 {
		maxAge = conf.MaxAge
	}
	opts := []handlers.CORSOption{
		handlers.AllowedOrigins(conf.AllowOrigins),
		handlers.AllowedMethods(conf.AllowMethods),
		handlers.AllowedHeaders(conf.AllowHeaders),
		handlers.ExposedHeaders(conf.ExposeHeaders),
		handlers.MaxAge(int(maxAge.Seconds())),
	}

	if conf.AllowCredentials {
		opts = append(opts, handlers.AllowCredentials())
	}

	return thttp.Filter(handlers.CORS(opts...))
}
