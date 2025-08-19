package request

import (
	"context"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"

	"github.com/limes-cloud/kratosx/config"
)

func TestRequest(t *testing.T) {
	conf := &config.Request{
		EnableLog:        false,
		RetryCount:       3,
		RetryWaitTime:    1 * time.Second,
		MaxRetryWaitTime: 1 * time.Second,
		Timeout:          1 * time.Second,
		UserAgent:        "test",
	}

	Init(conf, nil)

	request := Instance(context.Background())
	resp, err := request.Option(func(r *resty.Request) {
		r.Header.Set("test", "test")
	}).Get("https://www.baidu.com")
	assert.Nil(t, err)
	assert.NotNil(t, resp.Body())
}
