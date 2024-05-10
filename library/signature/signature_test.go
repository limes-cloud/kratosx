package signature

import (
	"testing"
	"time"

	"github.com/limes-cloud/kratosx/config"
)

func TestSign(t *testing.T) {
	tests := []struct {
		content string
		ak      string
		sk      string
		ts      int64
	}{
		{
			content: "123456",
		},
		{
			content: "12345678",
		},
	}

	watcher := func(key string, o config.WatchHandleFunc) {}
	Init(&config.Signature{
		Ak:     "app1",
		Sk:     "123456",
		Expire: 3 * time.Second,
	}, watcher)

	signer := Instance()
	for _, item := range tests {
		ts, sign, err := signer.Generate([]byte(item.content))
		if err != nil {
			t.Error(err.Error())
		}

		// 解密
		if err := signer.Verify([]byte(item.content), sign, ts); err != nil {
			t.Error(err.Error())
		}
	}
}
