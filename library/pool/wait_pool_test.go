package pool

import (
	"context"
	"testing"
	"time"

	"github.com/limes-cloud/kratosx/config"
)

func TestNewWaitRunner(t *testing.T) {
	Init(&config.Pool{Size: 1000}, func(key string, o config.WatchHandleFunc) {})

	wr := NewWaitRunner(WithWaitRunnerOption(1))
	start := time.Now().UnixMilli()
	for i := 0; i < 100; i++ {
		wr.AddTask(context.Background(), func() error {
			time.Sleep(10 * time.Millisecond)
			return nil
		})
	}
	end := time.Now().UnixMilli()

	wr.Wait()

	t.Log(end - start)
}
