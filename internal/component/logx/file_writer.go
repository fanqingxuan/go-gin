package logx

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type FileLevelWriter struct {
	Dirname     string
	FilePattern string
}

var _ zerolog.LevelWriter = (*FileLevelWriter)(nil)

func (w *FileLevelWriter) Write(p []byte) (n int, err error) {
	w.output("", p)
	return len(p), nil
}

func (w FileLevelWriter) WriteLevel(l zerolog.Level, p []byte) (n int, err error) {
	// s := string(p)
	suffix := ""
	switch l {
	case zerolog.WarnLevel, zerolog.ErrorLevel, zerolog.FatalLevel:
		suffix = "-error"
	}
	w.output(suffix, p)
	return len(p), nil
}

func (w FileLevelWriter) output(suffix string, p []byte) {
	pattern := time.Now().Format(w.FilePattern)
	dir := w.Dirname
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0744)
		if err != nil {
			panic(fmt.Errorf("can't make directories for new logfile: %s", err))
		}
	}

	filename := dir + pattern + suffix + ".log"
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("can not open or create file %s,error=%s", filename, err.Error()))
	}
	defer logFile.Close()
	logFile.Write(p)
}
