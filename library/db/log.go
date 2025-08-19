package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	glogger "gorm.io/gorm/logger"

	"github.com/limes-cloud/kratosx/library/logger"
)

// zap 适配gorm 日志
type sqlLog struct {
	logger        logger.Logger
	LogLevel      glogger.LogLevel
	SlowThreshold time.Duration
}

func newLog(level int, slow time.Duration) *sqlLog {
	return &sqlLog{
		logger:        logger.New(logger.WithCallerSkip(4)),
		LogLevel:      glogger.LogLevel(level),
		SlowThreshold: slow,
	}
}

func (l *sqlLog) Log(ctx context.Context) logger.Logger {
	return l.logger.WithContext(ctx)
}

func (l *sqlLog) LogMode(level glogger.LogLevel) glogger.Interface {
	l.LogLevel = level
	return l
}

func (l *sqlLog) Info(ctx context.Context, msg string, data ...any) {
	if l.LogLevel >= glogger.Info && len(data) > 0 {
		fields, _ := data[0].(map[string]any)
		l.Log(ctx).Info(msg, fields)
	}
}

func (l *sqlLog) Warn(ctx context.Context, msg string, data ...any) {
	if l.LogLevel >= glogger.Warn && len(data) > 0 {
		fields, _ := data[0].(map[string]any)
		l.Log(ctx).Warn(msg, fields)
	}
}

func (l *sqlLog) Error(ctx context.Context, msg string, data ...any) {
	if l.LogLevel >= glogger.Error && len(data) > 0 {
		fields, _ := data[0].(map[string]any)
		l.Log(ctx).Error(msg, fields)
	}
}

func (l *sqlLog) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= glogger.Silent {
		return
	}
	elapsed := time.Since(begin)
	costTime := float64(elapsed.Nanoseconds()) / 1e6
	switch {
	case err != nil && l.LogLevel >= glogger.Error && (!errors.Is(err, glogger.ErrRecordNotFound)):
		sql, rows := fc()
		l.Error(ctx, "sql错误", getSqlInfo(err.Error(), sql, rows, costTime, false))
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= glogger.Warn:
		sql, rows := fc()
		l.Warn(ctx, "sql告警", getSqlInfo("", sql, rows, costTime, true))
	case l.LogLevel == glogger.Info:
		sql, rows := fc()
		l.Info(ctx, "sql信息", getSqlInfo("", sql, rows, costTime, false))
	}
}

func getSqlInfo(err, sql string, rows int64, costTime float64, slow bool) map[string]any {
	return map[string]any{
		"err":  err,
		"sql":  sql,
		"rows": rows,
		"time": fmt.Sprintf("%vms", costTime),
		"slow": slow,
	}
}
