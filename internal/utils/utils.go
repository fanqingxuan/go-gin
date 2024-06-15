package utils

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

var ProjectDir string

func init() {
	_, file, _, _ := runtime.Caller(0)
	// compatible solution to get gorm source directory with various operating systems
	dir := filepath.Dir(filepath.Dir(filepath.Dir(file)))
	ProjectDir = filepath.ToSlash(dir) + "/"
	fmt.Println(ProjectDir)
}

// FileWithLineNum return the file name and line number of the current file
func FileWithLineNum() string {
	pcs := [13]uintptr{}
	// the third caller usually from gorm internal
	len := runtime.Callers(3, pcs[:])
	frames := runtime.CallersFrames(pcs[:len])
	for i := 0; i < len; i++ {
		// second return value is "more", not "ok"
		frame, _ := frames.Next()

		if (strings.HasPrefix(frame.File, ProjectDir) && !strings.HasSuffix(frame.File, "_test.go")) && !strings.HasSuffix(frame.File, ".gen.go") {
			return string(strconv.AppendInt(append([]byte(strings.ReplaceAll(frame.File, ProjectDir, "")), ':'), int64(frame.Line), 10))
		}
	}

	return ""
}
