package middleware

import (
	"context"

	"github.com/limes-cloud/kratosx/library/recovery"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
)

func Recovery() middleware.Middleware {
	handler := func(ctx context.Context, req, err any) error {
		e, ok := err.(*errors.Error)
		if ok {
			return e
		}
		return recovery.ErrUnknownRequest
	}

	return recovery.Recovery(recovery.WithHandler(handler))
}
