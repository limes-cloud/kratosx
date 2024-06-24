package captcha

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/mojocn/base64Captcha"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/email"
	"github.com/limes-cloud/kratosx/library/redis"
)

type Captcha interface {
	Email(tp string, ip string, to string) (Response, error)
	Image(tp string, ip string) (Response, error)
	VerifyEmail(tp, ip, id, answer, email string) error
	VerifyImage(tp, ip, id, answer string) error
}

type captcha struct {
	mu  sync.RWMutex
	set map[string]*config.Captcha
}

type Sender struct {
	UUID string
	Send func(conf *config.Captcha, answer string, expire time.Duration) (string, error)
}

var instance *captcha

const (
	imageType = "image"
	emailType = "email"
)

func Instance() Captcha {
	return instance
}

func Init(cfs map[string]*config.Captcha, watcher config.Watcher) {
	if len(cfs) == 0 {
		return
	}

	instance = &captcha{set: cfs}

	for key, conf := range cfs {
		if conf == nil {
			continue
		}

		instance.initFactory(key, conf)

		watcher("captcha."+key, func(value config.Value) {
			if err := value.Scan(conf); err != nil {
				log.Errorf("Captcha 配置变更失败：%s", err.Error())
				return
			}
			instance.initFactory(key, conf)
		})
	}
}

func (c *captcha) initFactory(name string, conf *config.Captcha) {
	c.mu.Lock()
	c.set[name] = conf
	c.mu.Unlock()
}

func (c *captcha) Image(tp string, ip string) (Response, error) {
	// 发送邮件
	sender := Sender{
		Send: func(conf *config.Captcha, answer string, expire time.Duration) (string, error) {
			// 生成验证码对应图片的base64
			dt := base64Captcha.NewDriverDigit(conf.Height, conf.Width, conf.Length, conf.Skew, conf.DotCount)
			item, err := dt.DrawCaptcha(answer)
			if err != nil {
				return "", err
			}
			return item.EncodeB64string(), err
		},
		UUID: "",
	}

	return c.generate(tp, ip, imageType, sender)
}

func (c *captcha) Email(tp string, ip string, to string) (Response, error) {
	// 发送邮件
	sender := Sender{
		Send: func(conf *config.Captcha, answer string, expire time.Duration) (string, error) {
			err := email.Instance().Template(conf.Template).Send(to, "", map[string]any{
				"captcha": answer,
				"minute":  int(conf.Expire.Minutes()),
			})
			return "", err
		},
		UUID: to,
	}

	return c.generate(tp, ip, emailType, sender)
}

func (c *captcha) generate(tp, ip, tpe string, sender Sender) (Response, error) {
	conf, is := c.set[tp]
	if !is {
		return nil, fmt.Errorf("%s captcha not exist", tp)
	}

	// 获取验证码存储器
	cache := redis.Instance().Get(conf.Redis)
	if cache == nil {
		return nil, fmt.Errorf("redis %v not exist", conf.Redis)
	}

	// 生成随机验证码
	answer := c.randomCode(conf.Length)

	// 获取当前用户的场景唯一id
	clientKey := c.clientUid(tp, ip, tpe)

	countKey := fmt.Sprintf("%s_count", clientKey)

	// 判断ip是否限制次数
	if conf.IpLimit != 0 {
		if count, _ := cache.Get(context.Background(), countKey).Int(); count > conf.IpLimit {
			return nil, errors.New("当前IP已超过最大验证次数")
		}
	}

	// 清除上一次生成的结果,防止同时造成大量生成请求占用内存
	if uid, _ := cache.Get(context.Background(), clientKey).Result(); uid != "" {
		if !conf.Refresh {
			return nil, errors.New("请勿重复请求验证码")
		}
		cache.Del(context.Background(), uid)
	}

	// 执行发送器
	base64, err := sender.Send(conf, answer, conf.Expire)
	if err != nil {
		return nil, err
	}

	// 获取当前验证码验证码唯一id
	uid := c.uid(clientKey, answer, sender.UUID)

	// 将本次验证码挂载到当前的场景id上
	if err := cache.Set(context.Background(), clientKey, uid, conf.Expire).Err(); err != nil {
		return nil, err
	}

	// 存储发送次数
	if conf.IpLimit != 0 && cache.Incr(context.Background(), countKey).Val() == 1 {
		// 设置当天00:00过期
		timeStr := time.Now().Format("2006-01-02")
		t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
		cache.ExpireAt(context.Background(), countKey, t.Add(86400*time.Second))
	}

	// 返回生成结果
	return &response{
		id:     uid,
		expire: conf.Expire,
		base64: base64,
	}, nil
}

func (c *captcha) VerifyEmail(tp, ip, id, answer, email string) error {
	return c.verify(tp, ip, emailType, id, answer, email)
}

func (c *captcha) VerifyImage(tp, ip, id, answer string) error {
	return c.verify(tp, ip, imageType, id, answer, "")
}

func (c *captcha) verify(tp, ip, name, id, answer, sender string) error {
	// 获取指定模板的配置
	conf, is := c.set[tp]
	if !is {
		return fmt.Errorf("%s captcha not exist", tp)
	}

	// 获取验证码存储器
	cache := redis.Instance().Get(conf.Redis)
	if cache == nil {
		return fmt.Errorf("redis %v not exist", conf.Redis)
	}

	// 获取当前用户的场景唯一id
	redisKey := c.clientUid(tp, ip, name)
	uid := c.uid(redisKey, answer, sender)

	// 获取用户当前的验证码场景id
	rid, err := cache.Get(context.Background(), redisKey).Result()
	if err != nil {
		return err
	}

	// 对比用户当前的验证码场景是否一致
	if rid != uid {
		return fmt.Errorf("captcha id %s  not exist", id)
	}

	// 验证通过清除缓存
	return cache.Del(context.Background(), redisKey).Err()
}

// randomCode 生成随机数验证码
func (c *captcha) randomCode(len int) string {
	rand.New(rand.NewSource(time.Now().Unix()))
	var code = rand.Intn(int(math.Pow10(len)) - int(math.Pow10(len-1)))
	return strconv.Itoa(code + int(math.Pow10(len-1)))
}

// uid 获取唯一id
func (c *captcha) clientUid(tp, ip, name string) string {
	return fmt.Sprintf("captcha:%s:%s:%x", name, tp, md5.Sum([]byte(ip)))
}

// uid 获取唯一id
func (c *captcha) uid(cid, ans, sender string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s:%s:%s", cid, ans, sender))))
}
