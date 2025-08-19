package logger

import (
	"context"
	"fmt"
	"os"
	"sync"

	kratoszap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/log"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/mbndr/figlet4go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/pkg"
)

var (
	// 全局日志器
	ins *logger

	// 构造器
	once sync.Once
)

type Field map[string]any

// F 构造日志字段
func F(key string, val any) Field {
	return Field{key: val}
}

type Logger interface {
	// Log 原始log方法
	Log(level log.Level, kvs ...any) error

	// Info 日志
	Info(msg string, fs ...Field)

	// Warn 日志
	Warn(msg string, fs ...Field)

	// Error 日志
	Error(msg string, fs ...Field)

	// Art 打印艺术字体
	Art(msg string)

	// WithContext 载入ctx
	WithContext(ctx context.Context) Logger

	// Sync 刷新日志
	Sync() error
}

type logger struct {
	zap    *zap.Logger
	opts   *options
	logger log.Logger
}

// Instance 获取全局日志器
func Instance() Logger {
	if ins == nil {
		return &logger{
			zap:    nil,
			opts:   &options{},
			logger: log.With(log.DefaultLogger),
		}
	}
	return ins
}

// New 复制全局日志器
func New(opts ...Option) Logger {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}

	if ins == nil {
		return &logger{
			zap:    nil,
			opts:   &options{},
			logger: log.With(log.DefaultLogger),
		}
	}

	// 复制日志器
	return &logger{
		zap:    ins.zap,
		opts:   o,
		logger: log.With(ins.logger, o.GetKV()...),
	}
}

// Init 初始化日志
func Init(conf *config.Logger, opts ...Option) {
	// 没配置则跳过
	if conf == nil {
		return
	}

	// 初始化日志器
	once.Do(func() {
		o := &options{}
		for _, opt := range opts {
			opt(o)
		}

		// 初始化
		zapLog := newZapLogger(conf)
		ins = &logger{
			zap:    zapLog,
			opts:   o,
			logger: log.With(kratoszap.NewLogger(zapLog), o.GetKV()...),
		}

		// 设置全局logger，格式化项目内置的打印器
		log.SetLogger(ins.logger)
	})
}

func (l *logger) Log(level log.Level, kvs ...any) error {
	return l.logger.Log(level, kvs...)
}

// WithContext 载入ctx
func (l *logger) WithContext(ctx context.Context) Logger {
	return &logger{
		zap:    l.zap,
		opts:   l.opts,
		logger: log.WithContext(ctx, l.logger),
	}
}

// Sync 刷新日志
func (l *logger) Sync() error {
	return l.zap.Sync()
}

// Info 日志
func (l *logger) Info(msg string, fs ...Field) {
	_ = l.logger.Log(log.LevelInfo, l.kvs(msg, fs...)...)
}

// Warn 日志
func (l *logger) Warn(msg string, fs ...Field) {
	_ = l.logger.Log(log.LevelWarn, l.kvs(msg, fs...)...)
}

// Error 日志
func (l *logger) Error(msg string, fs ...Field) {
	_ = l.logger.Log(log.LevelError, l.kvs(msg, fs...)...)
}

// Art 艺术打印输出
func (l *logger) Art(msg string) {
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	hexColor, _ := figlet4go.NewTrueColorFromHexString("885DBA")
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorGreen,
		figlet4go.ColorYellow,
		figlet4go.ColorCyan,
		hexColor,
	}

	renderStr, _ := ascii.RenderOpts(msg, options)
	fmt.Println(renderStr)
}

// kvs 格式化字段
func (l *logger) kvs(msg string, fs ...Field) []any {
	fs = append(fs, Field{"msg": msg})
	var kvs []any
	for _, f := range fs {
		for k, v := range f {
			kvs = append(kvs, k, v)
		}
	}
	return kvs
}

// newZapLogger 默认使用zap的日志器
func newZapLogger(conf *config.Logger) *zap.Logger {
	// 编码器配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "log",
		CallerKey:      "caller",
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,                          // 小写编码器
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"), // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 获取输出器
	getOutputs := func(conf *config.Logger, isError bool) []zapcore.WriteSyncer {
		var output []zapcore.WriteSyncer
		for _, val := range conf.Output {
			if val == "file" {
				fn := conf.File.Name
				if isError {
					fn = pkg.AppendFileSuffix(conf.File.Name, "_err")
				}
				filer, _ := rotatelogs.New(
					fn+".%Y%m%d%H",
					rotatelogs.WithLinkName(fn),
					rotatelogs.WithRotationTime(conf.File.SplitTime),
					rotatelogs.WithMaxAge(conf.File.MaxAge),
					rotatelogs.WithRotationCount(conf.File.MaxBackup),
				)
				output = append(output, zapcore.AddSync(filer))
			}

			if val == "stdout" {
				output = append(output, zapcore.AddSync(os.Stdout))
			}
		}
		return output
	}

	// // getStdOutputs 获取std输出器
	// hookStd := func(conf *config.Logger) {
	//	var output []zapcore.WriteSyncer
	//	// hookStdout
	//	fn := pkg.AppendFileSuffix(conf.File.Name, "_stdout")
	//	stdoutFile, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//	if err == nil {
	//		os.Stdout = stdoutFile
	//		output = append(output, zapcore.AddSync(stdoutFile))
	//	}
	//
	//	// hookStderr
	//	fn = pkg.AppendFileSuffix(conf.File.Name, "_stderr")
	//	stderrFile, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//	if err == nil {
	//		os.Stderr = stderrFile
	//		output = append(output, zapcore.AddSync(stderrFile))
	//	}
	//	return output
	// }

	// 编码器配置
	var encoder zapcore.Encoder
	if conf.EnCoder == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	var logs []zapcore.Core
	if conf.File != nil && conf.File.ErrorAlone {
		logs = []zapcore.Core{
			zapcore.NewCore(
				encoder,
				zapcore.NewMultiWriteSyncer(getOutputs(conf, false)...),
				zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
					return lvl < zapcore.WarnLevel
				}),
			),
			zapcore.NewCore(
				encoder,
				zapcore.NewMultiWriteSyncer(getOutputs(conf, true)...),
				zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
					return lvl >= zapcore.WarnLevel
				}),
			),
		}
	} else {
		logs = []zapcore.Core{
			zapcore.NewCore(
				encoder,
				zapcore.NewMultiWriteSyncer(getOutputs(conf, false)...),
				zapcore.Level(conf.Level),
			),
		}
	}

	// if conf.HookStd && conf.File != nil {
	//	logs = append(
	//		logs,
	//		zapcore.NewCore(
	//			encoder,
	//			zapcore.NewMultiWriteSyncer(getStdOutputs(conf)...),
	//			zapcore.Level(conf.Level),
	//		),
	//	)
	// }

	// 添加回调
	var zapOptions []zap.Option
	if conf.Caller {
		callerSkip := 3
		if conf.CallerSkip != 0 {
			callerSkip = conf.CallerSkip
		}
		zapOptions = append(zapOptions, zap.AddCaller())
		zapOptions = append(zapOptions, zap.AddCallerSkip(callerSkip))
	}

	return zap.New(zapcore.NewTee(logs...), zapOptions...)
}
