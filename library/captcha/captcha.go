package captcha

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type GetCaptchaRequest struct {
	Scene      string // 验证码场景
	ClientIP   string // 客户端IP地址
	User       string // 用户标识
	VerifyCode string // 设置验证码内容
}

type GetCaptchaResponse struct {
	UUID       string        // 验证码唯一标识
	Expire     time.Duration // 验证码有效期
	VerifyCode string        // 验证码内容
}

type VerifyCaptchaRequest struct {
	Scene      string // 验证码场景
	ClientIP   string // 客户端IP地址
	User       string // 用户标识
	UUID       string // 验证码唯一标识
	VerifyCode string // 验证码内容
}

// Captcha 验证码接口能力
type Captcha interface {
	// GetCaptchaDuration 获取验证码有效期
	GetCaptchaDuration() time.Duration

	// GetCaptcha 获取验证码
	GetCaptcha(req *GetCaptchaRequest) (*GetCaptchaResponse, error)

	// CancelCaptcha 取消验证码
	CancelCaptcha(uuid string) error

	// VerifyCaptcha 验证验证码(不同limit)
	VerifyCaptcha(req *VerifyCaptchaRequest) error

	// IsLimitError 判断是否超过最大获取验证码次数
	IsLimitError(err error) bool

	// IsDupError 判断是否重复请求验证码
	IsDupError(err error) bool

	// IsInvalidCaptchaError 判断是否验证码无效
	IsInvalidCaptchaError(err error) bool
}

type captcha struct {
	ctx            context.Context
	redis          *redis.Client
	limit          int           // 验证码在指定ip下的每日最大获取次数
	length         int           // 验证码长度
	expire         time.Duration // 验证码过期时间
	refreshTime    time.Duration // 刷新等待时间
	uniqueDevice   bool          // 是否唯一设备
	verifiedDelete bool          // 验证码验证成功后是否删除
}

var (
	// 这里两个error 上层调用方可能需要更具具体的error信息来处理业务。
	// 所以这里定义为全局变量，提供判断代码进行调用
	errorLimit          = errors.New("当前IP已超过最大获取验证码次数")
	errorDup            = errors.New("请勿重复请求验证码")
	errorInvalidCaptcha = errors.New("无效的验证码")
)

const (
	captchaDefaultLimit          = 50
	captchaDefaultLength         = 4
	captchaDefaultExpireTime     = 10 * time.Minute
	captchaDefaultRefreshTime    = 1 * time.Minute
	captchaDefaultUniqueDevice   = false
	captchaDefaultVerifiedDelete = false
)

type OptionFunc func(*captcha)

// WithLimit 设置每分钟最大获取次数
func WithLimit(limit int) OptionFunc {
	return func(c *captcha) {
		c.limit = limit
	}
}

// WithLength 设置验证码长度
func WithLength(length int) OptionFunc {
	return func(c *captcha) {
		c.length = length
	}
}

// WithExpire 设置验证码过期时间
func WithExpire(expire time.Duration) OptionFunc {
	return func(c *captcha) {
		c.expire = expire
	}
}

// WithUniqueDevice 设置是否唯一设备
func WithUniqueDevice(unique bool) OptionFunc {
	return func(c *captcha) {
		c.uniqueDevice = unique
	}
}

// WithRefresh 设置刷新验证码需要等待的时间
func WithRefresh(wait time.Duration) OptionFunc {
	return func(c *captcha) {
		c.refreshTime = wait
	}
}

// WithVerifiedDelete 验证成功后是否删除验证码
func WithVerifiedDelete(is bool) OptionFunc {
	return func(c *captcha) {
		c.verifiedDelete = is
	}
}

// WithContext 设置上下文
func WithContext(ctx context.Context) OptionFunc {
	return func(c *captcha) {
		c.ctx = ctx
	}
}

// NewCaptcha 初始化captcha对象
func NewCaptcha(redis *redis.Client, opts ...OptionFunc) Captcha {
	option := &captcha{
		redis:          redis,
		ctx:            context.Background(),
		limit:          captchaDefaultLimit,
		length:         captchaDefaultLength,
		expire:         captchaDefaultExpireTime,
		refreshTime:    captchaDefaultRefreshTime,
		uniqueDevice:   captchaDefaultUniqueDevice,
		verifiedDelete: captchaDefaultVerifiedDelete,
	}
	for _, opt := range opts {
		opt(option)
	}
	return option
}

// GetCaptchaDuration 获取验证码过期时间
func (c *captcha) GetCaptchaDuration() time.Duration {
	if c.refreshTime != 0 {
		return c.refreshTime
	}
	return c.expire
}

// VerifyCaptcha 验证验证码
func (c *captcha) VerifyCaptcha(req *VerifyCaptchaRequest) error {
	// 获取当前场景下客户端的唯一id
	sid := c.sid(req.Scene, req.ClientIP)

	// 通过uuid生成uid
	uid := c.getUIDByUUID(req.UUID)

	// 唯一设备校验
	if c.uniqueDevice && uid != c.uid(sid, req.User) {
		return errorInvalidCaptcha
	}

	// 获取验证码的唯一aid
	aid := c.aid(req.Scene, req.User, req.VerifyCode)

	// 判断验证码是否正确,或过期
	oriAid, err := c.redis.Get(c.ctx, uid).Result()
	if err != nil {
		return errorInvalidCaptcha
	}
	if oriAid != aid {
		return errorInvalidCaptcha
	}

	// 清除验证码
	if c.verifiedDelete {
		return c.redis.Del(c.ctx, uid).Err()
	}
	return nil
}

// incrLimitScript 原子检查并递增每日计数，返回递增后的值
// KEYS[1] = sid, ARGV[1] = limit, ARGV[2] = ttl(ms)
// 返回: 递增后的计数值
var incrLimitScript = redis.NewScript(`
local count = redis.call("INCR", KEYS[1])
if count == 1 then
	redis.call("PEXPIRE", KEYS[1], ARGV[2])
end
return count
`)

// GetCaptcha 获取验证码
func (c *captcha) GetCaptcha(req *GetCaptchaRequest) (*GetCaptchaResponse, error) {
	// 生成随机验证码
	verifyCode := c.randomCode(c.length)
	if req.VerifyCode != "" {
		verifyCode = req.VerifyCode
	}

	// 获取当前场景下客户端的唯一id
	sid := c.sid(req.Scene, req.ClientIP)

	// 原子判断并递增每日获取次数
	if c.limit != 0 {
		now := time.Now()
		ttl := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location()).Sub(now)
		count, err := incrLimitScript.Run(c.ctx, c.redis, []string{sid}, c.limit, ttl.Milliseconds()).Int()
		if err != nil {
			return nil, err
		}
		if count > c.limit {
			return nil, errorLimit
		}
	}

	// 获取当前场景下客户端用户的唯一id
	uid := c.uid(sid, req.User)

	// 判断是否允许刷新验证码
	if ttl := c.redis.TTL(c.ctx, uid).Val(); ttl.Seconds() > 0 {
		if c.refreshTime > 0 && (c.expire.Seconds()-ttl.Seconds()) < c.refreshTime.Seconds() {
			return nil, errorDup
		}
		// 清除旧的验证码
		c.redis.Del(c.ctx, uid)
	}

	// 生成验证码的唯一id
	aid := c.aid(req.Scene, req.User, verifyCode)

	// 设置验证码
	if err := c.redis.Set(c.ctx, uid, aid, c.expire).Err(); err != nil {
		return nil, err
	}


	return &GetCaptchaResponse{
		UUID:       c.getUUIDByUID(uid),
		VerifyCode: verifyCode,
		Expire:     c.expire,
	}, nil
}

// CancelCaptcha 取消验证码
func (c *captcha) CancelCaptcha(uuid string) error {
	// 获取当前场景下客户端用户的唯一id
	uid := c.getUIDByUUID(uuid)
	return c.redis.Del(c.ctx, uid).Err()
}

// IsLimitError 判断是否超过最大获取验证码次数
func (c *captcha) IsLimitError(err error) bool {
	return errors.Is(err, errorLimit) || errors.Is(err, errorDup)
}

// IsDupError 判断是否重复请求验证码
func (c *captcha) IsDupError(err error) bool {
	return errors.Is(err, errorDup)
}

// IsInvalidCaptchaError 判断是否无效的验证码
func (c *captcha) IsInvalidCaptchaError(err error) bool {
	return errors.Is(err, errorInvalidCaptcha)
}

// ClearDeviceLimit 清除设备限制
func (c *captcha) ClearDeviceLimit(scene, ip string) error {
	return c.redis.Del(c.ctx, c.sid(scene, ip)).Err()
}

// randomCode 生成随机数验证码
func (c *captcha) randomCode(len int) string {
	code := rand.Intn(int(math.Pow10(len)) - int(math.Pow10(len-1)))
	return strconv.Itoa(code + int(math.Pow10(len-1)))
}

// sid 生成当前场景下客户端的唯一id
func (c *captcha) sid(scene, ip string) string {
	return fmt.Sprintf("captcha:s:%x", md5.Sum([]byte(fmt.Sprintf("%s:%s", scene, ip))))
}

// uid 生成当前场景下客户端用户的唯一id
func (c *captcha) uid(sid, user string) string {
	return fmt.Sprintf("captcha:u:%x", md5.Sum([]byte(fmt.Sprintf("%s:%s", sid, user))))
}

// getUUIDByUID 从uid中提取uuid
func (c *captcha) getUUIDByUID(uid string) string {
	return strings.TrimPrefix(uid, "captcha:u:")
}

// getUidByUuid 生成验证码的唯一uid
func (c *captcha) getUIDByUUID(uuid string) string {
	return fmt.Sprintf("captcha:u:%s", uuid)
}

// aid 生成验证码答案的唯一标识
func (c *captcha) aid(scene, user, answer string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s:%s:%s", scene, user, answer))))
}
