package logx

import (
	"github.com/labstack/gommon/color"
	"github.com/rs/zerolog"
)

type ConsoleLevelWriter struct {
}

var _ zerolog.LevelWriter = (*ConsoleLevelWriter)(nil)

func (w *ConsoleLevelWriter) Write(p []byte) (n int, err error) {
	color.Print(color.White(string(p)))
	return len(p), nil
}

func (w ConsoleLevelWriter) WriteLevel(l zerolog.Level, p []byte) (n int, err error) {
	s := string(p)
	switch l {
	case zerolog.WarnLevel:
		s = color.Magenta(s)
	case zerolog.ErrorLevel:
		s = color.Red(s)
	case zerolog.FatalLevel:
		s = color.Bold(color.Red(s))
	default:
		s = color.White(s)
	}
	color.Print(s)
	return len(p), nil
}
