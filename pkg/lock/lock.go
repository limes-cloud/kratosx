package lock

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	goredis "github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/google/uuid"
	"time"
)

//https://cloud.tencent.com/developer/article/2334541

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
	AcquireFunc(f func() error, do func() error) error
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

func (l *lock) AcquireFunc(f func() error, do func() error) error {
	return nil
}
