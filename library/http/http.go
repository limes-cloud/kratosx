package http

import (
	"time"
	"unsafe"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-resty/resty/v2"
	json "github.com/json-iterator/go"

	"github.com/limes-cloud/kratosx/config"
)

type request struct {
	c        *config.Http
	request  *resty.Request
	logger   *log.Helper
	inputLog bool
}

type Request interface {
	DisableLog() Request
	Option(fn RequestFunc) Request
	Get(url string) (*response, error)
	Post(url string, data any) (*response, error)
	PostJson(url string, data any) (*response, error)
	Put(url string, data any) (*response, error)
	PutJson(url string, data any) (*response, error)
	Delete(url string) (*response, error)
}

func NewDefault(logger *log.Helper) Request {
	return New(&config.Http{
		EnableLog:        true,
		RetryCount:       3,
		RetryWaitTime:    100 * time.Millisecond,
		MaxRetryWaitTime: 3 * time.Second,
		Timeout:          10 * time.Second,
	}, logger)
}

func New(conf *config.Http, logger *log.Helper) Request {
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
	return &request{
		c:        conf,
		request:  client.R(),
		logger:   logger,
		inputLog: true,
	}
}

type RequestFunc func(*resty.Request) *resty.Request

func (h *request) DisableLog() Request {
	h.inputLog = false
	return h
}

func (h *request) Option(fn RequestFunc) Request {
	h.request = fn(h.request)
	return h
}

func (h *request) log(t int64, res *response) {
	if !(h.c.EnableLog && h.inputLog) {
		return
	}

	resData := res.Body()
	logs := []any{
		"type", "request",
		"method", h.request.Method,
		"url", h.request.URL,
		"header", h.request.Header,
		"body", h.request.Body,
		"cost", time.Now().UnixMilli() - t,
		"res", *(*string)(unsafe.Pointer(&resData)),
	}
	if len(h.request.FormData) != 0 {
		logs = append(logs, "form-data", h.request.FormData)
	}
	if len(h.request.QueryParam) != 0 {
		logs = append(logs, "query", h.request.QueryParam)
	}
	h.logger.Infow(logs...)
}

func (h *request) Get(url string) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.Get(url)
	return res, res.err
}

func (h *request) Post(url string, data any) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.SetBody(data).Post(url)
	return res, res.err
}

func (h *request) PostJson(url string, data any) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.ForceContentType("application/json").SetBody(data).Post(url)
	return res, res.err
}

func (h *request) Put(url string, data any) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.SetBody(data).Put(url)
	return res, res.err
}

func (h *request) PutJson(url string, data any) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.ForceContentType("application/json").SetBody(data).Put(url)
	return res, res.err
}

func (h *request) Delete(url string) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.Delete(url)
	return res, res.err
}

type response struct {
	err      error
	response *resty.Response
}

func (r *response) Body() []byte {
	return r.response.Body()
}

func (r *response) Result(val any) error {
	return json.Unmarshal(r.response.Body(), val)
}
