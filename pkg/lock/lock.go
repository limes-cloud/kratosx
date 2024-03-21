package lock

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/google/uuid"
)

type Option func(*option)

type option struct {
	timeout time.Duration
	value   string
}

func WithTimeout(timeout time.Duration) Option {
	return func(o *option) {
		o.timeout = timeout
	}
}

func WithValue(value string) Option {
	return func(o *option) {
		o.value = value
	}
}

type Lock interface {
	Acquire(ctx context.Context) error
	Release(ctx context.Context) error
	AcquireFunc(ctx context.Context, cf func() error, do func() error) error
}

type lock struct {
	mux    *redsync.Mutex
	option *option
}

func New(client *redis.Client, key string, opts ...Option) Lock {
	opt := &option{timeout: 8 * time.Second, value: uuid.NewString()}
	for _, of := range opts {
		of(opt)
	}

	pool := goredis.NewPool(client)
	return &lock{
		mux: redsync.New(pool).NewMutex(key,
			redsync.WithExpiry(opt.timeout),
			redsync.WithValue(opt.value),
		),
		option: opt,
	}
}

func (l *lock) Acquire(ctx context.Context) error {
	return l.mux.LockContext(ctx)
}

func (l *lock) Release(ctx context.Context) error {
	_, err := l.mux.UnlockContext(ctx)
	return err
}

func (l *lock) AcquireFunc(ctx context.Context, cf func() error, do func() error) error {
	for {
		// 获取数据，先查询，查询数据比插入数据qps更高，也防止锁变成串行化
		if err := cf(); err == nil {
			return nil
		}

		// 数据不存在则去拿锁
		if err := l.mux.TryLockContext(ctx); err == nil {
			break
		} else if !errors.Is(err, redsync.ErrFailed) {
			return err
		}

		// 防止频繁自旋[5-30ms]
		time.Sleep(time.Duration(5+rand.Intn(25)) * time.Millisecond)
	}

	// 查询缓存数据，防止获取锁之后，存在
	if err := cf(); err == nil {
		return nil
	}

	return func() error {
		defer l.mux.UnlockContext(ctx)
		return do()
	}()
}
