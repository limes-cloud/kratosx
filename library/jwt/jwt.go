package jwt

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	kratosJwt "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	json "github.com/json-iterator/go"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/redis"
)

type Jwt interface {
	NewToken(m map[string]any) (string, error)
	Parse(ctx context.Context, dst any) error
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
	instance *jwt
	tokenKey struct{}
)

const (
	blackPrefix  = "token_black"
	uniquePrefix = "token_unique"
)

// Instance 获取email对象实例
func Instance() Jwt {
	return instance
}

func Init(conf *config.JWT, watcher config.Watcher) {
	if conf == nil {
		return
	}

	instance = &jwt{conf: conf}

	watcher("jwt", func(value config.Value) {
		if err := value.Scan(conf); err != nil {
			log.Errorf("JWT 配置变更失败：%s", err.Error())
			return
		}

		instance.rw.Lock()
		defer instance.rw.Unlock()
		instance.conf = conf
	})
}

// NewToken is create jwt []byte
func (j *jwt) NewToken(m map[string]any) (string, error) {
	if j == nil {
		return "", errors.New("jwt config not enable or configure")
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

func (j *jwt) GetToken(ctx context.Context) string {
	token, _ := ctx.Value(tokenKey).(string)
	return token
}

func (j *jwt) SetToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenKey, token)
}

func (j *jwt) Renewal(ctx context.Context) (string, error) {
	token := j.GetToken(ctx)
	if token == "" {
		return "", errors.New("token is miss")
	}

	tokenInfo, _ := jwtv5.Parse(token, func(token *jwtv5.Token) (any, error) {
		return []byte(j.conf.Secret), nil
	})
	if tokenInfo == nil || tokenInfo.Claims == nil {
		return "", errors.New("token parse error")
	}

	claims, is := tokenInfo.Claims.(jwtv5.MapClaims)
	if !is {
		return "", errors.New("token format error")
	}

	// 判断token失效是否超过10s
	exp := int64(claims["exp"].(float64))
	now := time.Now().Unix()
	if exp > now {
		return "", errors.New("token is alive")
	}

	if now-exp > int64(j.conf.Renewal.Seconds()) {
		return "", errors.New("token is over max renewal time")
	}

	return j.NewToken(claims)
}

func (a *jwt) path(path, method string) string {
	return fmt.Sprintf("%s:%s", path, method)
}

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

func (j *jwt) IsBlacklist(token string) bool {
	rd := redis.Instance().Get(j.conf.Redis)
	if rd == nil {
		return false
	}
	is, _ := rd.HExists(context.Background(), blackPrefix, token).Result()
	return is
}

func (j *jwt) AddBlacklist(token string) {
	rd := redis.Instance().Get(j.conf.Redis)
	rd.HSet(context.Background(), blackPrefix, token, 1, j.conf.Expire)
}

func (j *jwt) setUnique(key, token string) {
	rd := redis.Instance().Get(j.conf.Redis)
	rd.HSet(context.Background(), uniquePrefix, key, token, j.conf.Expire)
}

func (j *jwt) CompareUniqueToken(key, token string) bool {
	rd := redis.Instance().Get(j.conf.Redis)
	res, _ := rd.HGet(context.Background(), uniquePrefix, key).Result()
	return res == token
}
