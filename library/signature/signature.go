package signature

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	kerrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	json "github.com/json-iterator/go"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/logger"
)

const (
	defaultTime = time.Second * 10

	timeHeader    = "x-md-sign-time"
	tokenHeader   = "x-md-sign-token"
	signReason    = "SignatureInvalid"
	genSignReason = "GenSignatureInvalid"
)

type Signature interface {
	// Generate 生成签名
	Generate(content []byte) (int64, string, error)

	// Verify 验证签名
	Verify(content []byte, sign string, ts int64) error

	// IsWhitelist 是否为白名单
	IsWhitelist(name string) bool

	// Server 服务端中间件
	Server() middleware.Middleware

	// Client 客户端中间件
	Client(conf *config.Signature) middleware.Middleware
}

type signature struct {
	conf *config.Signature
	mux  sync.RWMutex
}

var (
	ins *signature

	once sync.Once
)

func Instance() Signature {
	return ins
}

// Init 初始化签名
func Init(ec *config.Signature, watcher config.Watcher) {
	if ec == nil {
		return
	}

	once.Do(func() {
		if ec.Expire == 0 {
			ec.Expire = defaultTime
		}
		ins = &signature{
			conf: ec,
		}

		watcher("signature", func(value config.Value) {
			ins.mux.Lock()
			defer ins.mux.Unlock()

			if err := value.Scan(&ins.conf); err != nil {
				logger.Instance().Error("signature 配置变更失败", logger.F("err", err))
				return
			}
			if ins.conf.Expire == 0 {
				ins.conf.Expire = defaultTime
			}
		})
	})
}

// Generate 生成签名
func (s *signature) Generate(content []byte) (int64, string, error) {
	ts := time.Now().Unix()
	timestamp := strconv.FormatInt(ts, 10)
	// 添加时间戳
	content = append(content, []byte(fmt.Sprintf("|%s", timestamp))...)
	// 添加ak
	content = append(content, []byte(fmt.Sprintf("|%s", s.conf.Ak))...)
	// 加签
	her := hmac.New(sha256.New, []byte(s.conf.Sk))
	her.Write(content)
	return ts, hex.EncodeToString(her.Sum(nil)), nil
}

// Verify 验证签名
func (s *signature) Verify(content []byte, sign string, ts int64) error {
	// 解码
	sig, err := hex.DecodeString(sign)
	if err != nil {
		return err
	}

	if int64(s.conf.Expire.Seconds()) < (time.Now().Unix() - ts) {
		return errors.New("signature has expired")
	}

	timestamp := strconv.FormatInt(ts, 10)
	// 添加时间戳
	content = append(content, []byte(fmt.Sprintf("|%s", timestamp))...)
	// 添加ak
	content = append(content, []byte(fmt.Sprintf("|%s", s.conf.Ak))...)
	// 验签
	her := hmac.New(sha256.New, []byte(s.conf.Sk))
	her.Write(content)
	if !hmac.Equal(sig, her.Sum(nil)) {
		return errors.New("signature is invalid")
	}
	return nil
}

// IsWhitelist 判断是否为白名单
func (s *signature) IsWhitelist(name string) bool {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.conf.Whitelist[name]
}

// Server 服务端签名中间件
func (s *signature) Server() middleware.Middleware {
	if s == nil || !s.conf.Enable {
		return nil
	}
	return selector.Server(func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			md, ok := metadata.FromServerContext(ctx)
			if !ok {
				return handler(ctx, req)
			}

			timer := md.Get(timeHeader)
			sign := md.Get(tokenHeader)
			if timer == "" {
				return nil, kerrors.BadRequest(signReason, "must exist header:"+timeHeader)
			}
			ts, err := strconv.ParseInt(timer, 10, 64)
			if err != nil {
				return nil, kerrors.BadRequest(signReason, "time format error")
			}

			dataByte, err := json.Marshal(req)
			if err != nil {
				return nil, kerrors.BadRequest(signReason, err.Error())
			}

			if err := s.Verify(dataByte, sign, ts); err != nil {
				return nil, kerrors.BadRequest(signReason, err.Error())
			}
			return handler(ctx, req)
		}
	}).Match(func(ctx context.Context, operation string) bool {
		path := ""
		if h, is := http.RequestFromServerContext(ctx); is {
			path = h.Method + ":" + h.URL.Path
		}
		return !(s.IsWhitelist(operation) || s.IsWhitelist(path))
	}).Build()
}

// Client 客户端签名中间件
func (s *signature) Client(conf *config.Signature) middleware.Middleware {
	if conf == nil || !conf.Enable {
		return nil
	}
	cs := signature{
		conf: conf,
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			tr, ok := transport.FromClientContext(ctx)
			if !ok {
				return handler(ctx, req)
			}
			header := tr.RequestHeader()
			body, _ := json.Marshal(req)
			ts, token, err := cs.Generate(body)
			if err != nil {
				return nil, kerrors.BadRequest(genSignReason, err.Error())
			}

			header.Add(timeHeader, fmt.Sprint(ts))
			header.Add(tokenHeader, token)
			return handler(ctx, req)
		}
	}
}
