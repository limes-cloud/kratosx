package captcha

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	kratosConfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/email"
	"github.com/limes-cloud/kratosx/library/redis"
	"github.com/mojocn/base64Captcha"
	"math"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Captcha interface {
	Email(tp string, ip string, to string) (Response, error)
	Image(tp string, ip string) (Response, error)
	VerifyEmail(tp, ip, id, answer string) error
	VerifyImage(tp, ip, id, answer string) error
}

type captcha struct {
	mu  sync.RWMutex
	set map[string]*config.Captcha
}

var instance *captcha

const (
	imageType = "imageType"
	emailType = "emailType"
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

		watcher("captcha."+key, func(value kratosConfig.Value) {
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
	conf, is := c.set[tp]
	if !is {
		return nil, errors.New(fmt.Sprintf("%s captcha not exist", tp))
	}

	// 生成随机验证码
	answer := c.randomCode(conf.Length)

	// 生成验证码对应图片的base64
	dt := base64Captcha.NewDriverDigit(conf.Height, conf.Width, conf.Length, conf.Skew, conf.DotCount)
	item, err := dt.DrawCaptcha(answer)
	if err != nil {
		return nil, err
	}

	// 获取验证码存储器
	cache := redis.Instance().Get(conf.Redis)
	if cache == nil {
		return nil, fmt.Errorf("redis %v not exist", conf.Redis)
	}

	// 获取当前用户的场景唯一id
	redisKey := c.uid(tp, ip, imageType)

	// 清除上一次生成的结果,防止同时造成大量生成请求占用内存
	if uid, _ := cache.Get(context.Background(), redisKey).Result(); uid != "" {
		cache.Del(context.Background(), uid)
	}

	// 获取当前验证码验证码唯一id
	uid := uuid.New().String()
	if err = cache.Set(context.Background(), uid, answer, conf.Expire+time.Second).Err(); err != nil {
		return nil, err
	}

	// 将本次验证码挂载到当前的场景id上
	if err = cache.Set(context.Background(), redisKey, uid, conf.Expire+time.Second).Err(); err != nil {
		return nil, err
	}

	// 返回生成结果
	return &response{
		id:     uid,
		base64: item.EncodeB64string(),
		expire: conf.Expire,
	}, nil
}

func (c *captcha) Email(tp string, ip string, to string) (Response, error) {
	conf, is := c.set[tp]
	if !is {
		return nil, errors.New(fmt.Sprintf("%s captcha not exist", tp))
	}

	// 获取验证码存储器
	cache := redis.Instance().Get(conf.Redis)
	if cache == nil {
		return nil, fmt.Errorf("redis %v not exist", conf.Redis)
	}

	// 生成随机验证码
	answer := c.randomCode(conf.Length)

	// 获取当前用户的场景唯一id
	redisKey := c.uid(tp, ip, emailType)

	// 清除上一次生成的结果,防止同时造成大量生成请求占用内存
	if uid, _ := cache.Get(context.Background(), redisKey).Result(); uid != "" {
		cache.Del(context.Background(), uid)
	}

	// 获取当前验证码验证码唯一id
	uid := uuid.New().String()
	if err := cache.Set(context.Background(), uid, answer, conf.Expire+time.Second).Err(); err != nil {
		return nil, err
	}

	// 将本次验证码挂载到当前的场景id上
	if err := cache.Set(context.Background(), redisKey, uid, conf.Expire+time.Second).Err(); err != nil {
		return nil, err
	}

	if err := email.Instance().Template(conf.Template).Send(to, "", map[string]any{
		"answer": answer,
		"minute": int(conf.Expire.Minutes()),
	}); err != nil {
		return nil, err
	}
	// 返回生成结果
	return &response{
		id:     uid,
		expire: conf.Expire,
	}, nil
}

func (c *captcha) VerifyEmail(tp, ip, id, answer string) error {
	return c.verify(tp, ip, emailType, id, answer)
}

func (c *captcha) VerifyImage(tp, ip, id, answer string) error {
	return c.verify(tp, ip, imageType, id, answer)
}

func (c *captcha) verify(tp, ip, name, id, answer string) error {
	// 获取指定模板的配置
	conf, is := c.set[tp]
	if !is {
		return errors.New(fmt.Sprintf("%s captcha not exist", tp))
	}

	// 获取验证码存储器
	cache := redis.Instance().Get(conf.Redis)
	if cache == nil {
		return fmt.Errorf("redis %v not exist", conf.Redis)
	}

	// 获取当前用户的场景唯一id
	redisKey := c.uid(tp, ip, name)

	// 获取用户当前的验证码场景id
	rid, err := cache.Get(context.Background(), redisKey).Result()
	if err != nil {
		return err
	}

	// 对比用户当前的验证码场景是否一致
	if rid != id {
		return errors.New(fmt.Sprintf("captcha id %s  not exist", id))
	}

	// 获取指定验证码id的答案
	ans, err := cache.Get(context.Background(), id).Result()
	if err != nil {
		return err
	}
	// 对比答案是否一致
	if ans != answer {
		return errors.New("verify fail")
	}

	// 验证通过清除缓存
	return cache.Del(context.Background(), rid, id).Err()
}

// randomCode 生成随机数验证码
func (c *captcha) randomCode(len int) string {
	rand.New(rand.NewSource(time.Now().Unix()))
	var code = rand.Intn(int(math.Pow10(len)) - int(math.Pow10(len-1)))
	return strconv.Itoa(code + int(math.Pow10(len-1)))
}

// uid 获取唯一id
func (c *captcha) uid(tp, ip, name string) string {
	return fmt.Sprintf("captcha:%s:%s:%x", name, tp, md5.Sum([]byte(ip)))
}
