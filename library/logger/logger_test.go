package logger

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/limes-cloud/kratosx/config"
)

func TestInstance(t *testing.T) {
	conf := &config.Logger{
		Level:   0,
		Output:  []string{"stdout"},
		EnCoder: "json",
		Caller:  true,
		File: &config.LoggerFile{
			ErrorAlone: false,
			Name:       "test.log",
			MaxAge:     10,
			MaxBackup:  10,
		},
		CallerSkip: 1,
	}

	rewrite := func(name string) {
		file, err := os.Create(name)
		if err != nil {
			log.Fatal(err) // 这里用的是默认的stderr，因为它还没被重定向。可以先记录到控制台，然后重定向。
		}
		os.Stdout = file
	}

	pt := func(f func(Logger)) string {
		name := uuid.New().String() + ".txt"
		rewrite(name)
		Init(conf)

		// 打印
		logger := Instance()
		f(logger)
		_ = logger.Sync()
		defer os.Remove(name)

		text, _ := os.ReadFile(name)
		return string(text)
	}

	tests := []struct {
		input func() string
		want  string
	}{
		{
			input: func() string {
				return pt(func(logger Logger) {
					logger.Info("info", F("key", "value"))
				})
			},
			want: `{"level":"info","time":"2025-08-10 00:10:54.327","caller":"log/log.go:30","msg":"info","key":"value"}`,
		},
		{
			input: func() string {
				return pt(func(logger Logger) {
					logger.Warn("warn", F("key", "value"))
				})
			},
			want: `{"level":"warn","time":"2025-08-10 00:10:54.327","caller":"log/log.go:30","msg":"warn","key":"value"}`,
		},
		{
			input: func() string {
				return pt(func(logger Logger) {
					logger.Error("error", F("key", "value"))
				})
			},
			want: `{"level":"error","time":"2025-08-10 00:10:54.327","caller":"log/log.go:30","msg":"error","key":"value"}`,
		},
	}

	// 测试日志输出
	for _, test := range tests {
		wantMap := map[string]any{}
		_ = json.Unmarshal([]byte(test.want), &wantMap)
		wantMap["time"] = ""

		inputMap := map[string]any{}
		_ = json.Unmarshal([]byte(test.input()), &inputMap)
		inputMap["time"] = ""

		assert.Equal(t, wantMap, inputMap)
	}
}
