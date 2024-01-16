package middleware

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	ec "github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/signature"
	"strconv"
)

const (
	signReason   = "SignInvalid"
	formatReason = "MarshalRequestError"
	timeHeader   = "x-sign-time"
	signHeader   = "x-sign-token"
)

func Signature(conf *ec.Signature) middleware.Middleware {
	if conf == nil || !conf.Enable {
		return nil
	}
	return selector.Server(func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			header, ok := transport.FromServerContext(ctx)
			if !ok {
				return handler(ctx, req)
			}

			dataByte, err := json.Marshal(req)
			if err != nil {
				return nil, errors.BadRequest(formatReason, err.Error())
			}

			// 获取时间
			timeStr := header.RequestHeader().Get(timeHeader)
			if timeStr == "" {
				return nil, errors.BadRequest(formatReason, "must exist header:"+timeHeader)
			}
			ts, err := strconv.ParseInt(timeStr, 10, 64)
			if err != nil {
				return nil, errors.BadRequest(signReason, "time format error")
			}

			// 获取签名
			sign := header.RequestHeader().Get(signHeader)

			// 验签
			signIns := signature.Instance()
			if err := signIns.Verify(dataByte, sign, ts); err != nil {
				return nil, errors.BadRequest(signReason, signReason)
			}
			return handler(ctx, req)
		}
	}).Match(func(ctx context.Context, operation string) bool {
		path := ""
		if h, is := http.RequestFromServerContext(ctx); is {
			path = h.Method + ":" + h.URL.Path
		}
		signIns := signature.Instance()
		return !(signIns.IsWhitelist(operation) || signIns.IsWhitelist(path))
	}).Build()
}
