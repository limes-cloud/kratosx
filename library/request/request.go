package request

import (
	"context"
	"sync"
	"time"
	"unsafe"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-resty/resty/v2"
	json "github.com/json-iterator/go"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/logger"
)

type pr struct {
	conf   *config.Request
	ctx    context.Context
	logger logger.Logger
}

type request struct {
	c        *config.Request
	request  *resty.Request
	logger   logger.Logger
	inputLog bool
}

type Request interface {
	// DisableLog 禁用当前请求的日志
	DisableLog() Request

	// Option 设置请求数据
	Option(fn Func) Request

	// Get 发送get请求
	Get(url string) (*response, error)

	// Post 发送请求
	Post(url string, data any) (*response, error)

	// PostJson 发送json请求
	PostJson(url string, data any) (*response, error)

	// Put 发送json请求
	Put(url string, data any) (*response, error)

	// PutJson 发送json请求
	PutJson(url string, data any) (*response, error)

	// Delete 发送请求
	Delete(url string) (*response, error)

	// Do 执行请求
	Do() (*response, error)
}

var (
	// ins 请求配置
	ins *pr

	// once 初始化锁
	once = sync.Once{}
)

// Instance 获取实例
func Instance(ctx context.Context) Request {
	if ins == nil {
		once.Do(func() {
			ins = &pr{conf: &config.Request{EnableLog: true, RetryCount: 3, Timeout: 30 * time.Second}}
		})
	}
	tins := &pr{conf: ins.conf, ctx: ctx, logger: logger.Instance()}
	return tins.newRequest()
}

// Init 初始化
func Init(conf *config.Request, watcher config.Watcher) {
	once.Do(func() {
		ins = &pr{conf: conf}

		// 监听配置变更
		if watcher != nil {
			watcher("request", func(value config.Value) {
				if err := value.Scan(ins.conf); err != nil {
					log.Errorf("配置变更失败：%v", err.Error())
					return
				}
			})
		}
	})
}

// newRequest 创建请求
func (h *pr) newRequest() Request {
	conf := h.conf

	client := resty.New()
	if conf.MaxRetryWaitTime == 0 {
		conf.RetryWaitTime = 5 * time.Second
	}
	if conf.Timeout == 0 {
		conf.Timeout = 60 * time.Second
	}
	client.SetRetryWaitTime(conf.RetryWaitTime)
	client.SetRetryMaxWaitTime(conf.MaxRetryWaitTime)
	client.SetRetryCount(conf.RetryCount)
	client.SetTimeout(conf.Timeout)
	req := client.R()
	req = req.SetContext(h.ctx)
	if conf.UserAgent == "" {
		conf.UserAgent = "github.com/limes-cloud/kratosx http client"
	}
	req.Header.Set("User-Agent", conf.UserAgent)
	return &request{
		c:        conf,
		request:  req,
		logger:   h.logger,
		inputLog: true,
	}
}

type Func func(*resty.Request)

// DisableLog 禁用日志
func (h *request) DisableLog() Request {
	h.inputLog = false
	return h
}

// Option 执行可选参数
func (h *request) Option(fn Func) Request {
	fn(h.request)
	return h
}

// log 打印日志
func (h *request) log(t int64, res *response) {
	if !(h.c.EnableLog && h.inputLog) {
		return
	}

	resData := res.Body()
	logs := logger.Field{
		"index":  h.request.Attempt,
		"method": h.request.Method,
		"url":    h.request.URL,
		"header": h.request.Header,
		"body":   h.request.Body,
		"cost":   time.Now().UnixMilli() - t,
		"res":    *(*string)(unsafe.Pointer(&resData)),
	}
	if len(h.request.FormData) != 0 {
		logs["form-data"] = h.request.FormData
	}
	if len(h.request.QueryParam) != 0 {
		logs["query"] = h.request.QueryParam
	}
	h.logger.Info("request", logs)
}

// Get 发送get请求
func (h *request) Get(url string) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.Get(url)
	return res, res.err
}

// Post 发送post请求
func (h *request) Post(url string, data any) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.SetBody(data).Post(url)
	return res, res.err
}

// PostJson 发送post请求
func (h *request) PostJson(url string, data any) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.ForceContentType("application/json").SetBody(data).Post(url)
	return res, res.err
}

// Put 发送put请求
func (h *request) Put(url string, data any) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.SetBody(data).Put(url)
	return res, res.err
}

// PutJson 发送put请求
func (h *request) PutJson(url string, data any) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.ForceContentType("application/json").SetBody(data).Put(url)
	return res, res.err
}

// Delete 发送delete请求
func (h *request) Delete(url string) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.Delete(url)
	return res, res.err
}

// Do 执行请求
func (h *request) Do() (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.Send()
	return res, res.err
}

type response struct {
	err      error
	response *resty.Response
}

// Body 返回原始数据
func (r *response) Body() []byte {
	return r.response.Body()
}

// Result 解析返回结果
func (r *response) Result(val any) error {
	return json.Unmarshal(r.response.Body(), val)
}
