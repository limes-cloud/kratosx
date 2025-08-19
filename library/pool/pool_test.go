package pool

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/limes-cloud/kratosx/config"
)

func TestNewWaitRunner(t *testing.T) {
	conf := &config.Pool{
		Size:             100,
		ExpiryDuration:   10 * time.Second,
		PreAlloc:         true,
		MaxBlockingTasks: 100,
		Nonblocking:      false,
	}

	Init(conf)

	t.Run("test go", func(t *testing.T) {
		ch := make(chan int, 10)
		_ = Instance().GoFunc(func() {
			ch <- 1
		})

		assert.Equal(t, 1, <-ch)
	})

	t.Run("test context task", func(t *testing.T) {
		var (
			ch          = make(chan int, 10)
			ctx, cancel = context.WithCancel(context.Background())
		)

		cancel()
		_ = Instance().WithContext(ctx).GoFunc(func() {
			ch <- 1
		})

		time.Sleep(10 * time.Millisecond)
		close(ch)
		assert.NotEqual(t, 1, <-ch)
	})

	t.Run("test wait runner", func(t *testing.T) {
		count := 10
		runner := Instance().NewWaitRunner(
			WithMaxWaitRunnerOption(count),
			WithErrorBreakOption(true),
			WithRetryCountOption(3),
		)
		for i := 0; i < count; i++ {
			ind := i
			runner.AddTask(func() error {
				time.Sleep(1 * time.Second)
				if ind == count-1 {
					return fmt.Errorf("error index %d", ind)
				}
				return nil
			})
		}

		err := runner.Wait()
		assert.Error(t, err, "error index 9")
	})
}
