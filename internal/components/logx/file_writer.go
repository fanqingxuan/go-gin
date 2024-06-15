package logx

import (
	"github.com/labstack/gommon/color"
	"github.com/rs/zerolog"
)

type FileLevelWriter struct {
}

var _ zerolog.LevelWriter = (*FileLevelWriter)(nil)

func (w *FileLevelWriter) Write(p []byte) (n int, err error) {
	color.Print(color.White(string(p)))
	return len(p), nil
}

func (w FileLevelWriter) WriteLevel(l zerolog.Level, p []byte) (n int, err error) {
	// s := string(p)

	return len(p), nil
}
