package logger

import (
	"os"

	kratosZap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/limes-cloud/kratosx/config"
)

type LogField map[string]any

type logger struct {
	zap *zap.Logger
	fs  []any
}

var ins *logger

func Instance(opts ...Option) log.Logger {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	zapLog := ins.zap.WithOptions(zap.AddCallerSkip(o.callerSkip))
	return log.With(kratosZap.NewLogger(zapLog), ins.fs...)
}

func Helper(opts ...Option) *log.Helper {
	return log.NewHelper(Instance(opts...))
}

// Init 初始化日志器
func Init(lc *config.Logger, watcher config.Watcher, fields LogField) {
	// 没配置则跳过
	if lc == nil {
		return
	}

	// log field 转换
	var fs []any
	for key, val := range fields {
		fs = append(fs, key, val)
	}

	// 初始化
	ins = &logger{}
	ins.initFactory(lc, fs)

	watcher("log", func(value config.Value) {
		if err := value.Scan(lc); err != nil {
			log.Errorf("配置变更失败：%v", err.Error())
			return
		}
		// 变更初始化
		ins.initFactory(lc, fs)
	})
}

func (l *logger) initFactory(conf *config.Logger, fs []any) {
	// 创建zap logger
	l.zap = l.newZapLogger(conf)
	l.fs = fs

	gLog := log.With(kratosZap.NewLogger(l.zap), fs...)
	// 设置全局logger
	log.SetLogger(gLog)
}

func (l *logger) newZapLogger(conf *config.Logger) *zap.Logger {
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

	// 输出器配置
	var output []zapcore.WriteSyncer
	for _, val := range conf.Output {
		if val == "stdout" {
			output = append(output, zapcore.AddSync(os.Stdout))
		}
		if val == "file" {
			output = append(output, zapcore.AddSync(&lumberjack.Logger{
				Filename:   conf.File.Name,
				MaxSize:    conf.File.MaxSize,
				MaxBackups: conf.File.MaxBackup,
				MaxAge:     conf.File.MaxAge,
				Compress:   conf.File.Compress,
				LocalTime:  conf.File.LocalTime,
			}))
		}
	}

	var encoder zapcore.Encoder
	if conf.EnCoder == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}
	core := zapcore.NewCore(
		encoder,                                // 编码器配置
		zapcore.NewMultiWriteSyncer(output...), // 输出方式
		zapcore.Level(conf.Level),              // 设置日志级别
	)

	// 添加回调
	var zapOptions []zap.Option
	if conf.Caller {
		callerSkip := 3
		if conf.CallerSkip != nil {
			callerSkip = *conf.CallerSkip
		}
		zapOptions = append(zapOptions, zap.AddCaller())
		zapOptions = append(zapOptions, zap.AddCallerSkip(callerSkip))
	}

	return zap.New(core, zapOptions...)
}
