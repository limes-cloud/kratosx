package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"sync"
	"sync/atomic"
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
	km struct {
		keys  []K
		mu    sync.Mutex
		valid atomic.Bool
	}
	val       *sync.Map
	doMu      sync.Mutex
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

func (c *Cache[K, V]) onError(scene string, err error) {
	if c.hookerror != nil {
		c.hookerror(scene, err)
	} else {
		log.Printf("[cache] %s error: %v", scene, err)
	}
}

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
	msg.cache.doMu.Lock()
	defer msg.cache.doMu.Unlock()

	oldVal := msg.cache.copy()
	oldVs := msg.cache.vs

	defer func() {
		if err != nil {
			msg.cache.val = oldVal
			msg.cache.vs = oldVs
			msg.cache.km.valid.Store(false)
		}
	}()

	// 执行操作列表
	if err = msg.do(false); err != nil {
		return err
	}

	// 广播变更
	if err = msg.broadcast(); err != nil {
		return err
	}

	return nil
}

// do 执行操作链，localOnly 为 true 时仅更新本地缓存
func (msg *Msg[K, V]) do(localOnly bool) error {
	msg.cache.km.valid.Store(false)

	if !localOnly {
		pipe := msg.cache.rd.Pipeline()
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
		return msg.cache.setVersion(msg.ctx, pipe, msg.cache.version())
	}

	for _, item := range msg.op.List {
		switch item.Action {
		case opPut:
			msg.cache.store(item.Key, *item.Val)
		case opDel:
			msg.cache.delete(item.Key)
		}
	}
	msg.cache.vs = msg.cache.version()
	return nil
}

func (c *Cache[K, V]) transCacheKey(key K) string {
	return fmt.Sprint(key)
}

func (c *Cache[K, V]) transCacheVal(val V) string {
	cv := cacheValue[V]{
		Val: val,
	}
	b, _ := json.Marshal(cv)
	return string(b)
}

func (c *Cache[K, V]) Keys() []K {
	if c.km.valid.Load() {
		return c.km.keys
	}

	c.km.mu.Lock()
	defer c.km.mu.Unlock()

	if c.km.valid.Load() {
		return c.km.keys
	}

	var keys []K
	c.val.Range(func(key, _ any) bool {
		keys = append(keys, key.(K))
		return true
	})
	c.km.keys = keys
	c.km.valid.Store(true)
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
	src := c.Keys()
	keys := make([]K, len(src))
	copy(keys, src)
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
		select {
		case <-ctx.Done():
			return
		default:
		}

		m, err := c.queue.ReceiveMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return
			}
			c.onError("subscribe", err)
			continue
		}

		recvMsg := op[K, V]{}
		if err := json.Unmarshal([]byte(m.Payload), &recvMsg); err != nil {
			c.onError("unmarshal subscribe", err)
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
		c.doMu.Lock()
		_ = oper.do(true)
		c.doMu.Unlock()
	}
}

// Init 初始化数据
func (c *Cache[K, V]) Init(ctx context.Context) error {
	if c.hookload == nil {
		return fmt.Errorf("hookload is not configured")
	}
	ms, err := c.hookload()
	if err != nil {
		return err
	}

	for k, v := range ms {
		c.val.Store(k, v)
	}
	c.km.valid.Store(false)

	// 获取当前的版本
	rvs, _ := c.rd.Get(ctx, c.versionKey()).Result()
	vs := c.version()
	if vs != rvs {
		pipe := c.rd.Pipeline()
		pipe.Del(ctx, c.cacheKey())
		for k, v := range ms {
			pipe.HSet(ctx, c.cacheKey(), c.transCacheKey(k), c.transCacheVal(v))
		}
		if err := c.setVersion(ctx, pipe, c.version()); err != nil {
			return err
		}
	}

	return nil
}

// Repair 定时修复缓存，避免变更监听失败
func (c *Cache[K, V]) Repair(ctx context.Context) {
	ticker := time.NewTicker(c.wt)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}

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
		if c.hookload == nil {
			continue
		}
		m, err := c.hookload()
		if err != nil {
			c.onError("repair load", err)
			continue
		}

		c.doMu.Lock()

		// 替换本地缓存：先删除旧key，再写入新数据
		c.val.Range(func(key, _ any) bool {
			k := key.(K)
			if _, ok := m[k]; !ok {
				c.val.Delete(k)
			}
			return true
		})
		for k, v := range m {
			c.val.Store(k, v)
		}
		c.km.valid.Store(false)

		// 同步到 Redis
		pipe := c.rd.Pipeline()
		pipe.Del(ctx, c.cacheKey())
		for k, v := range m {
			pipe.HSet(ctx, c.cacheKey(), c.transCacheKey(k), c.transCacheVal(v))
		}
		vs := c.version()
		c.vs = vs
		pipe.Set(ctx, c.versionKey(), vs, redis.KeepTTL)
		if _, err := pipe.Exec(ctx); err != nil {
			c.onError("repair sync", err)
		}
		c.doMu.Unlock()
	}
}
