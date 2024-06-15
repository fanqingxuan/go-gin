package redisx

import (
	"context"
	"go-gin/internal/components/logx"
	"net"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
)

var sourceDir string

func init() {
	_, file, _, _ := runtime.Caller(0)
	// compatible solution to get gorm source directory with various operating systems
	dir := filepath.Dir(filepath.Dir(filepath.Dir(file)))
	sourceDir = filepath.ToSlash(dir) + "/"
}

type LogHook struct{}

func (LogHook) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		conn, err := next(ctx, network, addr)
		if err != nil {
			logx.WithContext(ctx).Warnf("redis", "dail error=%s", err.Error())
		}
		return conn, err
	}
}
func (LogHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		err := next(ctx, cmd)
		if err != nil {
			logx.WithContext(ctx).Warnf("redis", "%s execute command:%+v, error=%s", fileWithLineNum(), cmd.Args(), err)
		}
		return nil
	}
}
func (LogHook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		return next(ctx, cmds)
	}
}

// FileWithLineNum return the file name and line number of the current file
func fileWithLineNum() string {
	pcs := [13]uintptr{}
	// the third caller usually from gorm internal
	len := runtime.Callers(3, pcs[:])
	frames := runtime.CallersFrames(pcs[:len])
	for i := 0; i < len; i++ {
		// second return value is "more", not "ok"
		frame, _ := frames.Next()

		if (strings.HasPrefix(frame.File, sourceDir) && !strings.HasSuffix(frame.File, "_test.go")) && !strings.HasSuffix(frame.File, ".gen.go") {
			return string(strconv.AppendInt(append([]byte(strings.ReplaceAll(frame.File, sourceDir, "")), ':'), int64(frame.Line), 10))
		}
	}

	return ""
}
