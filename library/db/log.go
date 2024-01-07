package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/logger"

	libLogger "github.com/limes-cloud/kratosx/library/logger"
)

// zap 适配gorm 日志
type sqlLog struct {
	logger        *log.Helper
	LogLevel      logger.LogLevel
	SlowThreshold time.Duration
}

func newLog(level int, slow time.Duration) logger.Interface {
	return &sqlLog{
		logger:        libLogger.Helper(libLogger.AddCallerSkip(4)),
		LogLevel:      logger.LogLevel(level),
		SlowThreshold: slow,
	}
}

func (l *sqlLog) Log(ctx context.Context) *log.Helper {
	return l.logger.WithContext(ctx)
}

func (l *sqlLog) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *sqlLog) Info(ctx context.Context, msg string, data ...any) {
	if l.LogLevel >= logger.Info {
		data = append(data, "title", msg)
		l.Log(ctx).Infow(data...)
	}
}

func (l *sqlLog) Warn(ctx context.Context, msg string, data ...any) {
	if l.LogLevel >= logger.Warn {
		data = append(data, "title", msg)
		l.Log(ctx).Infow(data...)
	}
}

func (l *sqlLog) Error(ctx context.Context, msg string, data ...any) {
	if l.LogLevel >= logger.Error {
		data = append(data, "title", msg)
		l.Log(ctx).Infow(data...)
	}
}

func (l *sqlLog) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	costTime := float64(elapsed.Nanoseconds()) / 1e6
	switch {
	case err != nil && l.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound)):
		sql, rows := fc()
		l.Error(ctx, "sql错误", getSqlInfo(err.Error(), sql, rows, costTime, false)...)
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		l.Warn(ctx, "sql告警", getSqlInfo("", sql, rows, costTime, true)...)
	case l.LogLevel == logger.Info:
		sql, rows := fc()
		l.Info(ctx, "sql信息", getSqlInfo("", sql, rows, costTime, false)...)
	}
}

func getSqlInfo(err, sql string, rows int64, costTime float64, slow bool) []any {
	return []any{
		"err", err,
		"sql", sql,
		"rows", rows,
		"time", fmt.Sprintf("%vms", costTime),
		"slow", slow,
	}
}
