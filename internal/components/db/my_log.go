package db

import (
	"context"
	"errors"
	"fmt"
	"go-gin/internal/components/logx"
	"go-gin/internal/utils"
	"time"

	"gorm.io/gorm/logger"
)

var (
	infoStr      = "%s"
	warnStr      = "%s"
	errStr       = "%s "
	traceStr     = "%s [%.3fms] [rows:%v] %s"
	traceWarnStr = "%s %s [%.3fms] [rows:%v] %s"
	traceErrStr  = "%s %s [%.3fms] [rows:%v] %s"
)

type my_log struct {
}

// LogMode log mode
func (l *my_log) LogMode(level logger.LogLevel) logger.Interface {

	return &my_log{}
}

// Info print info
func (l *my_log) Info(ctx context.Context, msg string, data ...interface{}) {
	logx.WithContext(ctx).Infof("sql", infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
}

// Warn print warn messages
func (l *my_log) Warn(ctx context.Context, msg string, data ...interface{}) {
	logx.WithContext(ctx).Warnf("sql", warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
}

// Error print error messages
func (l *my_log) Error(ctx context.Context, msg string, data ...interface{}) {
	logx.WithContext(ctx).Errorf("sql", errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
}

// Trace print sql message
//
//nolint:cyclop
func (l *my_log) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

	elapsed := time.Since(begin)
	slowThreshold := 3 * time.Second
	switch {
	case err != nil && !errors.Is(err, logger.ErrRecordNotFound):
		sql, rows := fc()
		if rows == -1 {
			logx.WithContext(ctx).Errorf("sql", traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			logx.WithContext(ctx).Errorf("sql", traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > slowThreshold:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", slowThreshold)
		if rows == -1 {
			logx.WithContext(ctx).Debugf("sql", traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, '-', sql)
		} else {
			logx.WithContext(ctx).Debugf("sql", traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)

		}
	default:
		sql, rows := fc()
		if rows == -1 {
			logx.WithContext(ctx).Debugf("sql", traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			logx.WithContext(ctx).Debugf("sql", traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
