package middleware

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
)

func RateLimit(enable bool) middleware.Middleware {
	if !enable {
		return nil
	}
	return ratelimit.Server()
}
