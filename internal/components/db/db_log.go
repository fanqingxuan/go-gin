package db

import (
	"context"
	"errors"
	"fmt"
	"go-gin/internal/components/logx"
	"go-gin/internal/utils"
	"strings"
	"time"

	"gorm.io/gorm/logger"
)

var (
	traceStr     = "%s [%.3fms] [rows:%v] %s"
	traceWarnStr = "%s %s [%.3fms] [rows:%v] %s"
	traceErrStr  = "%s %s [%.3fms] [rows:%v] %s"
)

type db_log struct {
	LogLevel logger.LogLevel
}

func ParseLevel(levelStr string) logger.LogLevel {
	level_str := strings.ToLower(levelStr)
	switch level_str {
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
func (l *db_log) LogMode(level logger.LogLevel) logger.Interface {

	return &db_log{
		LogLevel: level,
	}
}

// Info print info
func (l *db_log) Info(ctx context.Context, msg string, data ...interface{}) {
}

// Warn print warn messages
func (l *db_log) Warn(ctx context.Context, msg string, data ...interface{}) {
}

// Error print error messages
func (l *db_log) Error(ctx context.Context, msg string, data ...interface{}) {
}

// Trace print sql message
//
//nolint:cyclop
func (l *db_log) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	slowThreshold := 3 * time.Second
	switch {
	case err != nil && l.LogLevel >= logger.Error && !errors.Is(err, logger.ErrRecordNotFound):
		sql, rows := fc()
		if rows == -1 {
			logx.WithContext(ctx).Errorf("sql", traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			logx.WithContext(ctx).Errorf("sql", traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > slowThreshold && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", slowThreshold)
		if rows == -1 {
			logx.WithContext(ctx).Debugf("sql", traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, '-', sql)
		} else {
			logx.WithContext(ctx).Debugf("sql", traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)

		}
	case l.LogLevel == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			logx.WithContext(ctx).Debugf("sql", traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			logx.WithContext(ctx).Debugf("sql", traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
