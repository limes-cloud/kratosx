package logger

import (
	"os"

	kratosZap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/limes-cloud/kratosx/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LogField map[string]any

type logger struct {
	log log.Logger
}

var instance *logger

func Instance() log.Logger {
	if instance == nil {
		return log.GetLogger()
	}
	return instance.log
}

func Helper() *log.Helper {
	return log.NewHelper(instance.log)
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
	instance = &logger{}
	instance.initFactory(lc, fs)

	watcher("log", func(value config.Value) {
		if err := value.Scan(lc); err != nil {
			log.Errorf("配置变更失败：%v", err.Error())
			return
		}
		// 变更初始化
		instance.initFactory(lc, fs)
	})
}

func (l *logger) initFactory(conf *config.Logger, fs []any) {
	// 创建zap logger
	l.log = log.With(l.newZapLogger(conf), fs...)
	// 设置全局logger
	log.SetLogger(instance.log)
}

func (l *logger) newZapLogger(conf *config.Logger) *kratosZap.Logger {
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

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),  // 编码器配置
		zapcore.NewMultiWriteSyncer(output...), // 输出方式
		zapcore.Level(conf.Level),              // 设置日志级别
	)

	return kratosZap.NewLogger(zap.New(core))
}
