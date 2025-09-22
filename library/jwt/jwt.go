package jwt

import (
	"context"
	"fmt"
	"regexp"
	"sync"
	"time"

	"github.com/limes-cloud/kratosx/library/logger"

	kerros "github.com/go-kratos/kratos/v2/errors"
	kratosJwt "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	json "github.com/json-iterator/go"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/redis"
)

const reason = "UNAUTHORIZED"

var (
	ErrInvalidTokenFormat     = kerros.Unauthorized(reason, "Invalid token format")
	ErrTokenParseFail         = kerros.Unauthorized(reason, "Failed to parse token")
	ErrJWTConfigNotConfigured = kerros.Unauthorized(reason, "JWT configuration is not enabled or not set")
	ErrTokenMaxRenewal        = kerros.Unauthorized(reason, "Token has exceeded the maximum renewal time")
	ErrTokenStillValid        = kerros.Unauthorized(reason, "Token is still valid")
)

type Jwt interface {
	NewToken(m map[string]any) (string, error)
	Parse(ctx context.Context, dst any) error
	ParseByToken(token string, dst any) error
	ParseMapClaims(ctx context.Context) (map[string]any, error)
	IsWhitelist(path, method string) bool
	IsBlacklist(token string) bool
	AddBlacklist(token string)
	GetToken(ctx context.Context) string
	SetToken(ctx context.Context, token string) context.Context
	Renewal(ctx context.Context) (string, error)
	CompareUniqueToken(key, token string) bool
}

type jwt struct {
	conf *config.JWT
	rw   sync.RWMutex
}

var (
	ins *jwt

	once sync.Once
)

type tokenKey struct{}

const (
	blackPrefix  = "token_black"
	uniquePrefix = "token_unique"
)

// Instance 获取email对象实例
func Instance() Jwt {
	return ins
}

// Init 初始化
func Init(conf *config.JWT, watcher config.Watcher) {
	if conf == nil {
		return
	}

	once.Do(func() {
		ins = &jwt{conf: conf}

		watcher("jwt", func(value config.Value) {
			ins.rw.Lock()
			defer ins.rw.Unlock()

			if err := value.Scan(&ins.rw); err != nil {
				logger.Instance().Info("jwt config watch error", logger.F("err", err))
				return
			}
		})
	})
}

// NewToken is create jwt []byte
func (j *jwt) NewToken(m map[string]any) (string, error) {
	if j == nil {
		return "", ErrJWTConfigNotConfigured
	}

	m["exp"] = jwtv5.NewNumericDate(time.Now().Add(j.conf.Expire + time.Second)) // 过期时间
	m["nbf"] = jwtv5.NewNumericDate(time.Now())                                  // 生效时间
	m["iat"] = jwtv5.NewNumericDate(time.Now())                                  // 签发时间

	keyFunc := func(token *jwtv5.Token) (any, error) {
		return []byte(j.conf.Secret), nil
	}

	jwtToken := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, jwtv5.MapClaims(m))
	key, err := keyFunc(jwtToken)
	if err != nil {
		return "", err
	}

	token, err := jwtToken.SignedString(key)
	if err != nil {
		return "", err
	}

	if j.conf.Unique {
		uniqueKey, ok := m[j.conf.UniqueKey]
		if ok {
			j.setUnique(fmt.Sprint(uniqueKey), token)
		}
	}

	return token, nil
}

// Parse 解析token
func (j *jwt) Parse(ctx context.Context, dst any) error {
	claims, err := j.ParseMapClaims(ctx)
	if err != nil {
		return err
	}
	body, err := json.Marshal(claims)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, dst)
}

// ParseByToken 解析指定的token
func (j *jwt) ParseByToken(token string, dst any) error {
	tokenInfo, _ := jwtv5.Parse(token, func(token *jwtv5.Token) (any, error) {
		return []byte(j.conf.Secret), nil
	})
	if tokenInfo == nil || tokenInfo.Claims == nil {
		return ErrTokenParseFail
	}

	claims, is := tokenInfo.Claims.(jwtv5.MapClaims)
	if !is {
		return ErrInvalidTokenFormat
	}

	body, err := json.Marshal(claims)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, dst)
}

// ParseMapClaims 解析token到map
func (j *jwt) ParseMapClaims(ctx context.Context) (map[string]any, error) {
	tokenInfo, is := kratosJwt.FromContext(ctx)
	if !is {
		if j == nil {
			return nil, kratosJwt.ErrMissingJwtToken
		}
		token := j.GetToken(ctx)
		if token == "" {
			return nil, kratosJwt.ErrMissingJwtToken
		}

		parser, _ := jwtv5.Parse(token, func(token *jwtv5.Token) (any, error) {
			return []byte(j.conf.Secret), nil
		})
		if parser == nil || parser.Claims == nil {
			return nil, kratosJwt.ErrTokenInvalid
		}

		tokenInfo, is = parser.Claims.(jwtv5.MapClaims)
		if !is {
			return nil, kratosJwt.ErrTokenParseFail
		}
	}
	claims, is := tokenInfo.(jwtv5.MapClaims)
	if !is {
		return nil, kratosJwt.ErrTokenParseFail
	}
	return claims, nil
}

// GetToken 从ctx中获取token
func (j *jwt) GetToken(ctx context.Context) string {
	token, _ := ctx.Value(tokenKey{}).(string)
	return token
}

// SetToken 设置token的值到当前的ctx
func (j *jwt) SetToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenKey{}, token)
}

// Renewal token续期
func (j *jwt) Renewal(ctx context.Context) (string, error) {
	token := j.GetToken(ctx)
	if token == "" {
		return "", kratosJwt.ErrMissingJwtToken
	}

	tokenInfo, _ := jwtv5.Parse(token, func(token *jwtv5.Token) (any, error) {
		return []byte(j.conf.Secret), nil
	})
	if tokenInfo == nil || tokenInfo.Claims == nil {
		return "", ErrTokenParseFail
	}

	claims, is := tokenInfo.Claims.(jwtv5.MapClaims)
	if !is {
		return "", ErrInvalidTokenFormat
	}

	// 判断token失效是否超过10s
	exp := int64(claims["exp"].(float64))
	now := time.Now().Unix()
	if exp > now {
		return "", ErrTokenStillValid
	}

	if now-exp > int64(j.conf.Renewal.Seconds()) {
		return "", ErrTokenMaxRenewal
	}

	return j.NewToken(claims)
}

func (a *jwt) path(path, method string) string {
	return fmt.Sprintf("%s:%s", path, method)
}

// IsWhitelist 判断请求的接口是否在白名单内
func (j *jwt) IsWhitelist(path, method string) bool {
	j.rw.RLock()
	defer j.rw.RUnlock()

	if !j.conf.EnableGrpc && method == "GRPC" {
		return true
	}

	path = j.path(method, path)
	if j.conf.Whitelist[path] {
		return true
	}

	for p := range j.conf.Whitelist {
		// 将*替换为匹配任意多字符的正则表达式
		pattern := "^" + p + "$"
		pattern = regexp.MustCompile(`\*`).ReplaceAllString(pattern, ".+")

		// 编译正则表达式
		re := regexp.MustCompile(pattern)

		// 检查输入是否匹配正则表达式
		if re.MatchString(path) {
			return true
		}
	}
	return false
}

// IsBlacklist 判断token是否在黑名单
func (j *jwt) IsBlacklist(token string) bool {
	rd := redis.Instance().Get(j.conf.Redis)
	if rd == nil {
		return false
	}
	is, _ := rd.HExists(context.Background(), blackPrefix, token).Result()
	return is
}

// AddBlacklist 添加token进入黑名单
func (j *jwt) AddBlacklist(token string) {
	rd := redis.Instance().Get(j.conf.Redis)
	rd.HSet(context.Background(), blackPrefix, token, 1, j.conf.Expire)
}

// setUnique 设置当前token为unique token
func (j *jwt) setUnique(key, token string) {
	rd := redis.Instance().Get(j.conf.Redis)
	rd.HSet(context.Background(), uniquePrefix, key, token, j.conf.Expire)
}

// CompareUniqueToken 对比是否时unique key
func (j *jwt) CompareUniqueToken(key, token string) bool {
	rd := redis.Instance().Get(j.conf.Redis)
	res, _ := rd.HGet(context.Background(), uniquePrefix, key).Result()
	return res == token
}
