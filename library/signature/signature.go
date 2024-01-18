package signature

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/limes-cloud/kratosx/config"
	"strconv"
	"sync"
	"time"
)

type Signature interface {
	// Generate 生成签名
	Generate(content []byte) (int64, string, error)
	// Verify 验证签名
	Verify(content []byte, sign string, ts int64) error
	// IsWhitelist 是否为白名单
	IsWhitelist(name string) bool
}

type signature struct {
	conf *config.Signature
	mu   sync.RWMutex
}

const defaultTime = time.Second * 3

var instance *signature

func Instance() Signature {
	return instance
}

func Init(ec *config.Signature, watcher config.Watcher) {
	if ec == nil {
		return
	}

	if ec.Time == 0 {
		ec.Time = defaultTime
	}
	instance = &signature{
		conf: ec,
	}

	watcher("signature", func(value config.Value) {
		nec := config.Signature{}
		if err := value.Scan(&nec); err != nil {
			log.Errorf("Signature 配置变更失败：%s", err.Error())
			return
		}
		if nec.Time == 0 {
			nec.Time = defaultTime
		}

		instance.mu.Lock()
		*instance.conf = nec
		instance.mu.Unlock()
	})
}

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

func (s *signature) Verify(content []byte, sign string, ts int64) error {
	//解码
	sig, err := hex.DecodeString(sign)
	if err != nil {
		return err
	}

	if int64(s.conf.Time.Seconds()) < (time.Now().Unix() - ts) {
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

func (s *signature) IsWhitelist(name string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.conf.Whitelist[name]
}
