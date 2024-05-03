package config

import (
	kratosConfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/protobuf/proto"
)

type WatchHandleFunc func(Value)

type Watcher func(key string, o WatchHandleFunc)

type Config interface {
	SetAppInfo(id, name, version string)
	Load() error
	Scan(v any) error
	Value(key string) Value
	Watch(key string, o WatchHandleFunc)
	ScanWatch(key string, o WatchHandleFunc)
	Close() error
	App() *App
	IsInit() bool
}

type config struct {
	app *App
	ins kratosConfig.Config
}

var instance *config

func Instance() Config {
	return instance
}

func New(source kratosConfig.Source) Config {
	instance = &config{
		ins: kratosConfig.New(
			kratosConfig.WithSource(source),
		),
	}
	return instance
}

func (c *config) IsInit() bool {
	return !(c == nil || c.ins == nil)
}

func (c *config) Load() error {
	if err := c.ins.Load(); err != nil {
		return err
	}
	c.app = new(App)
	return c.Scan(c.app)
}

func (c *config) App() *App {
	return c.app
}

func (c *config) Scan(dst any) error {
	if _, ok := dst.(proto.Message); ok {
		return c.ins.Scan(&dst)
	}

	// 序列化json
	res := map[string]any{}
	if err := c.ins.Scan(&res); err != nil {
		return err
	}

	dc := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           dst,
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
		),
	}
	decoder, err := mapstructure.NewDecoder(dc)
	if err != nil {
		return err
	}
	return decoder.Decode(res)
}

func (c *config) transformValue(val kratosConfig.Value) Value {
	return &value{Value: val}
}

func (c *config) Value(key string) Value {
	return c.transformValue(c.ins.Value(key))
}

func (c *config) ScanWatch(key string, handler WatchHandleFunc) {
	handler(c.Value(key))
	c.Watch(key, handler)
}

func (c *config) Watch(key string, handler WatchHandleFunc) {
	if err := c.ins.Watch(key, func(_ string, value kratosConfig.Value) {
		defer func() {
			if p := recover(); p != nil {
				log.Error("监听配置失败：%v", p)
			}
		}()
		handler(c.transformValue(value))
	}); err != nil {
		log.Error(key+"监听配置失败", err.Error())
	}
}

func (c *config) Close() error {
	return c.ins.Close()
}

func (c *config) SetAppInfo(id, name, version string) {
	c.app.ID, c.app.Name, c.app.Version = id, name, version
}
