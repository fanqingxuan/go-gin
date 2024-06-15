package middlewares

import (
	"go-gin/internal/components/logx"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var sourceDir string

func init() {
	_, file, _, _ := runtime.Caller(0)
	// compatible solution to get gorm source directory with various operating systems
	dir := filepath.Dir(filepath.Dir(filepath.Dir(file)))
	sourceDir = filepath.ToSlash(dir) + "/"
}

func recoverLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logx.PanicLoggerInstance.Fatal().
					Ctx(ctx).
					Any("error", err).
					Str("files", fileWithLineNum()).
					Send()
				ctx.Abort()
			}
		}()
		ctx.Next()

	}
}

// FileWithLineNum return the file name and line number of the current file
func fileWithLineNum() string {
	pcs := [13]uintptr{}
	// the third caller usually from gorm internal
	len := runtime.Callers(3, pcs[:])
	frames := runtime.CallersFrames(pcs[:len])
	filelist := []string{}
	for i := 0; i < len; i++ {
		// second return value is "more", not "ok"
		frame, _ := frames.Next()

		if (strings.HasPrefix(frame.File, sourceDir) && !strings.HasSuffix(frame.File, "_test.go")) && !strings.HasSuffix(frame.File, ".gen.go") {
			filelist = append(filelist, string(strconv.AppendInt(append([]byte(strings.ReplaceAll(frame.File, sourceDir, "")), ':'), int64(frame.Line), 10)))
		}
	}

	return strings.Join(filelist, "\r\n")
}
