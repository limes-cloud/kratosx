package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"sort"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"github.com/redis/go-redis/v9"
)

type KT interface {
	~string | ~int8 | ~int16 | ~int | ~int32 | ~int64 |
		~uint8 | ~uint16 | ~uint | ~uint32 | ~uint64
}

type Cache[K KT, V any] struct {
	val       *sync.Map
	node      string
	vs        string
	key       string
	wt        time.Duration
	rd        *redis.Client
	queue     *redis.PubSub
	hookerror func(scene string, err error)
	hookload  func() (map[K]V, error)
}

type cacheValue[V any] struct {
	Val V `json:"val"`
}

const (
	opPut = "put"
	opDel = "del"
)

type Msg[K KT, V any] struct {
	ctx   context.Context
	mux   *sync.Mutex
	cache *Cache[K, V]
	op    *op[K, V]
}

type op[K KT, V any] struct {
	Sender string          `json:"sender"` // 发送者
	List   []*opItem[K, V] `json:"list"`
}

type opItem[K KT, V any] struct {
	Action string `json:"action"` // 指定的动作
	Key    K      `json:"key"`
	Val    *V     `json:"val"`
}

type Option[K KT, V any] func(*Cache[K, V])

func HookLoad[K KT, V any](fn func() (map[K]V, error)) Option[K, V] {
	return func(c *Cache[K, V]) {
		c.hookload = fn
	}
}

func NewCache[K KT, V any](client *redis.Client, key string, opts ...Option[K, V]) *Cache[K, V] {
	cache := &Cache[K, V]{
		key:  key,
		val:  &sync.Map{},
		wt:   10 * time.Second,
		node: uuid.New().String(),
		rd:   client,
	}

	for _, opt := range opts {
		opt(cache)
	}

	cache.queue = client.Subscribe(context.Background(), cache.queueKey())

	return cache
}

func NewCacheAndInit[K KT, V any](ctx context.Context, client *redis.Client, key string, opts ...Option[K, V]) *Cache[K, V] {
	lc := NewCache[K, V](client, key, opts...)
	// 加载缓存,加载失败则直接报错，避免线上隐式错误。
	if err := lc.Init(ctx); err != nil {
		panic(err)
	}

	// 监听缓存变更
	go lc.Subscribe(ctx)

	// 定期修复缓存
	go lc.Repair(ctx)

	return lc
}

func (msg *Msg[K, V]) addOp(action string, key K, val *V) {
	msg.mux.Lock()
	defer msg.mux.Unlock()

	msg.op.List = append(msg.op.List, &opItem[K, V]{
		Action: action,
		Key:    key,
		Val:    val,
	})
}

func (msg *Msg[K, V]) Put(key K, value V) *Msg[K, V] {
	msg.addOp(opPut, key, &value)
	return msg
}

func (msg *Msg[K, V]) Puts(ms map[K]V) *Msg[K, V] {
	for k, v := range ms {
		msg.addOp(opPut, k, &v)
	}
	return msg
}

func (msg *Msg[K, V]) Delete(key K) *Msg[K, V] {
	msg.addOp(opDel, key, nil)
	return msg
}

func (msg *Msg[K, V]) Deletes(keys []K) *Msg[K, V] {
	for _, k := range keys {
		msg.addOp(opDel, k, nil)
	}
	return msg
}

func (msg *Msg[K, V]) Do() (err error) {
	oldVal := msg.cache.copy()
	oldVs := msg.cache.vs

	defer func() {
		if err != nil {
			msg.cache.val = oldVal
			msg.cache.vs = oldVs
		}
	}()

	// 执行操作列表
	if err = msg.do(); err != nil {
		return err
	}

	// 广播变更
	if err = msg.broadcast(); err != nil {
		return err
	}

	return nil
}

// do 执行操作链
func (msg *Msg[K, V]) do() error {
	pipe := msg.cache.rd.Pipeline()
	// 变更本地缓存
	for _, item := range msg.op.List {
		switch item.Action {
		case opPut:
			pipe.HSet(msg.ctx, msg.cache.cacheKey(), msg.cache.transCacheKey(item.Key), msg.cache.transCacheVal(*item.Val))
			msg.cache.store(item.Key, *item.Val)
		case opDel:
			pipe.HDel(msg.ctx, msg.cache.cacheKey(), msg.cache.transCacheKey(item.Key))
			msg.cache.delete(item.Key)
		}
	}

	// 重新计算版本
	return msg.cache.setVersion(msg.ctx, pipe, msg.cache.version())
}

func (c *Cache[K, V]) oriCacheKey(key string) K {
	var orival any = new(K)
	switch orival.(type) {
	case int8:
		return K(cast.ToInt8(key))
	case int16:
		return K(cast.ToInt16(key))
	case int32:
		return K(cast.ToInt32(key))
	case int:
		return K(cast.ToInt(key))
	case int64:
		return K(cast.ToInt64(key))
	case uint8:
		return K(cast.ToUint8(key))
	case uint16:
		return K(cast.ToUint16(key))
	case uint32:
		return K(cast.ToUint32(key))
	case uint:
		return K(cast.ToUint(key))
	case uint64:
		return K(cast.ToUint64(key))
	default:
		return orival.(K)
	}
}

func (c *Cache[K, V]) transCacheKey(key K) string {
	return fmt.Sprint(key)
}

func (c *Cache[K, V]) oriCacheVal(val string) V {
	var cv cacheValue[V]
	err := json.Unmarshal([]byte(val), &cv)
	if err != nil {
		return cv.Val
	}
	return cv.Val
}

func (c *Cache[K, V]) transCacheVal(val V) string {
	var cv = cacheValue[V]{
		Val: val,
	}
	b, _ := json.Marshal(cv)
	return string(b)
}

func (c *Cache[K, V]) Keys() []K {
	var keys []K
	c.val.Range(func(key, _ any) bool {
		keys = append(keys, key.(K))
		return true
	})
	return keys
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	v, ok := c.val.Load(key)
	if !ok {
		var empty V
		return empty, false
	}
	return v.(V), ok
}

func (c *Cache[K, V]) copy() *sync.Map {
	cm := &sync.Map{}
	c.val.Range(func(key, value any) bool {
		cm.Store(key, value)
		return true
	})
	return cm
}

func (c *Cache[K, V]) OP(ctx context.Context) *Msg[K, V] {
	return &Msg[K, V]{
		ctx:   ctx,
		cache: c,
		mux:   &sync.Mutex{},
		op: &op[K, V]{
			Sender: c.node,
			List:   make([]*opItem[K, V], 0),
		},
	}
}

func (c *Cache[K, V]) Version() string {
	return c.vs
}

// version 计算当前节点数据的版本号
func (c *Cache[K, V]) version() string {
	keys := c.Keys()
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	// k-v 对排列
	var arr []any
	for _, key := range keys {
		val, ok := c.val.Load(key)
		if !ok {
			continue
		}
		arr = append(arr, key, val)
	}

	// 计算md5值
	str, _ := json.Marshal(arr)
	return crypto.MD5(str)
}

func (c *Cache[K, V]) cacheKey() string {
	return fmt.Sprintf("cache:%s", c.key)
}

func (c *Cache[K, V]) queueKey() string {
	return fmt.Sprintf("cache:%s:queue", c.key)
}

func (c *Cache[K, V]) versionKey() string {
	return fmt.Sprintf("cache:%s:version", c.key)
}

func (c *Cache[K, V]) store(key K, val V) {
	c.val.Store(key, val)
}

func (c *Cache[K, V]) delete(key K) {
	c.val.Delete(key)
}

func (c *Cache[K, V]) setVersion(ctx context.Context, tx redis.Pipeliner, vs string) error {
	c.vs = vs
	version, _ := c.rd.Get(ctx, c.versionKey()).Result()
	if version != vs {
		tx.Set(ctx, c.versionKey(), vs, redis.KeepTTL)
		_, err := tx.Exec(ctx)
		return err
	}
	return nil
}

// broadcast 广播变更
func (msg *Msg[K, V]) broadcast() error {
	strMsg, _ := json.Marshal(msg.op)
	return msg.cache.rd.Publish(msg.ctx, msg.cache.queueKey(), string(strMsg)).Err()
}

// Subscribe 监听变更
func (c *Cache[K, V]) Subscribe(ctx context.Context) {
	for {
		m, err := c.queue.ReceiveMessage(ctx)
		if err != nil {
			c.hookerror("subscribe", err)
			continue
		}

		recvMsg := op[K, V]{}
		if err := json.Unmarshal([]byte(m.Payload), &recvMsg); err != nil {
			c.hookerror("unmarshal subscribe", err)
			continue
		}

		// 如果是当前节点，直接跳过
		if recvMsg.Sender == c.node {
			continue
		}

		oper := c.OP(ctx)
		for _, item := range recvMsg.List {
			oper.addOp(item.Action, item.Key, item.Val)
		}
		_ = oper.do()
	}
}

// Init 初始化数据
func (c *Cache[K, V]) Init(ctx context.Context) error {
	ms, err := c.hookload()
	if err != nil {
		return err
	}
	pipe := c.rd.Pipeline()
	for k, v := range ms {
		c.val.Store(k, v)
		pipe.HSet(context.Background(), c.cacheKey(), c.transCacheKey(k), c.transCacheVal(v))
	}
	if err := c.setVersion(ctx, pipe, c.version()); err != nil {
		return err
	}
	return nil
}

// Repair 定时修复缓存，避免变更监听失败
func (c *Cache[K, V]) Repair(ctx context.Context) {
	for {
		time.Sleep(c.wt)
		// 查询当前的版本号
		version, err := c.rd.Get(ctx, c.versionKey()).Result()
		if err != nil {
			continue
		}

		// 计算当前的版本，如果数据一致，则不修复
		if c.version() == version {
			continue
		}

		// 重新加载缓存
		m, err := c.hookload()
		if err != nil {
			c.hookerror("repair load", err)
		}
		for k, v := range m {
			c.val.Store(k, v)
		}
	}
}
