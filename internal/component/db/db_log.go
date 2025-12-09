package db

import (
	"context"
	"errors"
	"fmt"
	"go-gin/internal/component/logx"
	"go-gin/internal/util"
	"strings"
	"time"

	"gorm.io/gorm/logger"
)

var (
	traceStr     = "%s [%.3fms] [rows:%v] %s"
	traceWarnStr = "%s %s [%.3fms] [rows:%v] %s"
	traceErrStr  = "%s %s [%.3fms] [rows:%v] %s"
)

type DBLog struct {
	LogLevel logger.LogLevel
}

func ParseLevel(levelStr string) logger.LogLevel {
	l := strings.ToLower(levelStr)
	switch l {
	case "debug":
		return logger.Info
	case "info":
		return logger.Info
	case "warn":
		return logger.Warn
	case "error":
		return logger.Error
	}
	return logger.Info
}

// LogMode log mode
func (l *DBLog) LogMode(level logger.LogLevel) logger.Interface {

	return &DBLog{
		LogLevel: level,
	}
}

// Info print info
func (l *DBLog) Info(ctx context.Context, msg string, data ...any) {
}

// Warn print warn messages
func (l *DBLog) Warn(ctx context.Context, msg string, data ...any) {
}

// Error print error messages
func (l *DBLog) Error(ctx context.Context, msg string, data ...any) {
}

// Trace print sql message
//
//nolint:cyclop
func (l *DBLog) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	slowThreshold := 10 * time.Second
	switch {
	case err != nil && l.LogLevel >= logger.Error && !errors.Is(err, logger.ErrRecordNotFound):
		sql, rows := fc()
		if rows == -1 {
			logx.WithContext(ctx).Errorf("sql", traceErrStr, util.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			logx.WithContext(ctx).Errorf("sql", traceErrStr, util.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > slowThreshold && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", slowThreshold)
		if rows == -1 {
			logx.WithContext(ctx).Warnf("sql", traceWarnStr, util.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, '-', sql)
		} else {
			logx.WithContext(ctx).Warnf("sql", traceWarnStr, util.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)

		}
	case l.LogLevel == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			logx.WithContext(ctx).Debugf("sql", traceStr, util.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			logx.WithContext(ctx).Debugf("sql", traceStr, util.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
