package core

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/limes-cloud/kratosx/library/tasker"

	"github.com/google/uuid"
	"github.com/limes-cloud/kratosx/library/logger"
	lredis "github.com/limes-cloud/kratosx/library/redis"
	"github.com/redis/go-redis/v9"
)

type LockOption func(*lockOption)

type lockOption struct {
	acquireWait time.Duration
	timeout     time.Duration
	autoRenewal bool
	value       string
	redis       string
}

// WithLockTimeout 锁的超时时间
func WithLockTimeout(timeout time.Duration) LockOption {
	return func(o *lockOption) {
		o.timeout = timeout
	}
}

// WithLockValue 设置锁的值
func WithLockValue(value string) LockOption {
	return func(o *lockOption) {
		o.value = value
	}
}

// WithLockAcquireWaitTime 尝试获取锁的自旋等待时间
func WithLockAcquireWaitTime(timeout time.Duration) LockOption {
	return func(o *lockOption) {
		o.acquireWait = timeout
	}
}

// WithLockAutoRenewal 自动续期锁，如果当前持有锁则可以续期
func WithLockAutoRenewal() LockOption {
	return func(o *lockOption) {
		o.autoRenewal = true
	}
}

type Lock interface {
	// Acquire 获取不到则阻塞，直到获取到锁
	Acquire() error

	// TryAcquire 尝试获取锁，获取不到则返回false
	TryAcquire() (bool, error)

	// Release 释放锁
	Release() error

	// Renewal 续期锁，如果当前持有锁则可以续期
	Renewal() error

	// AcquireFunc 尝试获取锁，并执行do函数，执行完成后释放锁
	// 常用于查询缓存，如果不存在则查询数据库，并写入缓存
	// search 一般是用来查询redis 是否存在缓存数据，具体更具业务情况编写
	// do 一般是用来查询数据库，并把数据写入缓存
	AcquireFunc(search func() error, do func() error) error
}

const (
	// 续期
	renewalLuaScript = `
if redis.call("GET", KEYS[1]) == ARGV[1] then
	redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
    return "OK"
else
    return "no permission"
end`

	// 加锁
	lockLuaScript = `
if redis.call("GET", KEYS[1]) == ARGV[1] then
    return "OK"
else
    return redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2])
end`

	// 释放锁
	delLuaScript = `
if redis.call("GET", KEYS[1]) == ARGV[1] then
    return redis.call("DEL", KEYS[1])
else
    return 0
end
`
)

type lock struct {
	ctx     context.Context
	key     string
	store   *redis.Client
	option  *lockOption
	renewed atomic.Bool
	closed  atomic.Bool
}

// NewLock 初始化redis-lock
func NewLock(ctx context.Context, key string, opts ...LockOption) Lock {
	opt := &lockOption{timeout: 30 * time.Second, value: uuid.NewString()}
	for _, of := range opts {
		of(opt)
	}

	var (
		rd       = lredis.Instance()
		rdClient = rd.Get()
	)
	if opt.redis != "" {
		rdClient = rd.Get(opt.redis)
	}

	task := &lock{
		ctx:    ctx,
		key:    "lock:" + key,
		store:  rdClient,
		option: opt,
	}
	// 在服务关闭时释放锁，防止死锁
	tasker.Instance().BeforeStop("release redis lock", func() {
		_ = task.Release()
	})
	return task
}

// Acquire 获取锁，直到获取到为止
func (lc *lock) Acquire() error {
	for {
		// 获得锁
		unlock, err := lc.TryAcquire()
		if err != nil {
			return err
		}
		if unlock {
			return nil
		}
		// 防止频繁自旋[5-30ms]
		if lc.option.acquireWait > 0 {
			time.Sleep(lc.option.acquireWait)
		} else {
			time.Sleep(time.Duration(5+rand.Intn(25)) * time.Millisecond)
		}
	}
}

// TryAcquire 尝试获取锁，直接返回获取结果，不阻塞，允许重入
func (lc *lock) TryAcquire() (bool, error) {
	if lc.closed.Load() {
		return false, errors.New("lock is closed")
	}
	resp, err := lc.store.Do(lc.ctx,
		"eval",
		lockLuaScript,
		1,
		lc.key,
		lc.option.value,
		strconv.Itoa(int(lc.option.timeout.Milliseconds())),
	).Result()
	if err != nil && err.Error() != "redis: nil" {
		return false, fmt.Errorf("acquire lock error %s", err.Error())
	}
	if resp == nil || err != nil {
		return false, nil
	}
	reply, ok := resp.(string)
	if ok && reply == "OK" {
		// 自动续期
		if lc.option.autoRenewal && !lc.renewed.Load() {
			lc.autoRenewal()
		}
		return true, nil
	}
	return false, fmt.Errorf("acquire lock error %s", resp)
}

// autoRenewal 	自动续期锁，如果当前持有锁则可以续期
func (lc *lock) autoRenewal() {
	lc.renewed.Store(true)
	go func() {
		count := 0
		for {
			time.Sleep(lc.option.timeout - lc.option.timeout/3)
			if lc.closed.Load() {
				return
			}

			count++
			// 续期失败超过3次，则不再续期
			if err := lc.Renewal(); err != nil {
				if count > 3 {
					logger.Instance().WithContext(lc.ctx).Error("lock renewal error", logger.F("error", err))
					return
				}
				// 续期失败，等待100ms后重试，防止因为网络抖动导致续期失败
				time.Sleep(100 * time.Millisecond)
				continue
			}
			count = 0
		}
	}()
}

// Release 释放锁
func (lc *lock) Release() error {
	if lc.closed.Load() {
		return nil
	}

	resp, err := lc.store.Do(lc.ctx,
		"eval",
		delLuaScript,
		1,
		lc.key,
		lc.option.value,
	).Result()
	if err != nil {
		return err
	}

	reply, _ := resp.(int64)
	if reply != 1 {
		return errors.New("lock not released as expected")
	}

	lc.closed.Store(true)
	return nil
}

// AcquireFunc 尝试获取锁，并执行do函数，执行完成后释放锁
// 常用于查询缓存，如果不存在则查询数据库，并写入缓存
// search 一般是用来查询redis 是否存在缓存数据，具体更具业务情况编写
// do 一般是用来查询数据库，并把数据写入缓存
func (lc *lock) AcquireFunc(search func() error, do func() error) error {
	for {
		// 获取数据，先查询，查询数据比插入数据qps更高，也防止锁变成串行化
		if err := search(); err == nil {
			return nil
		}

		// 数据不存在则去拿锁
		unlock, err := lc.TryAcquire()
		if err != nil {
			return err
		}

		if unlock {
			break
		}

		// 防止频繁自旋[5-30ms]
		if lc.option.acquireWait > 0 {
			time.Sleep(lc.option.acquireWait)
		} else {
			time.Sleep(time.Duration(5+rand.Intn(25)) * time.Millisecond)
		}
	}

	// 查询缓存数据，防止获取锁之后，存在
	if err := search(); err == nil {
		return nil
	}

	return func() error {
		err := do()
		_ = lc.Release()
		return err
	}()
}

// Renewal 续租锁
func (lc *lock) Renewal() error {
	resp, err := lc.store.Do(lc.ctx,
		"eval",
		renewalLuaScript,
		1,
		lc.key,
		lc.option.value,
		strconv.Itoa(int(lc.option.timeout.Milliseconds())),
	).Result()
	if err != nil {
		return err
	}

	reply, ok := resp.(string)
	if ok && reply == "OK" {
		return nil
	}
	return fmt.Errorf("acquire lock error %s", resp)
}
