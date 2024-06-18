package logx

import (
	"io"
	"strings"
	"time"

	"github.com/labstack/gommon/color"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	ConsoleWriter = &ConsoleLevelWriter{}

	FileWriter = zerolog.SyncWriter(&FileLevelWriter{
		Dirname:     "",
		FilePattern: time.DateOnly,
	})

	AcessFileWriter = zerolog.SyncWriter(&FileLevelWriter{
		Dirname:     "acecss/",
		FilePattern: time.DateOnly,
	})
)

var (
	AccessLoggerInstance zerolog.Logger = zerolog.New(AcessFileWriter).Level(zerolog.InfoLevel).With().Timestamp().Logger().Hook(TracingHook{})
	PanicLoggerInstance                 = (zerolog.Logger{})
)

func Init(levelstr string, isDebugMode bool) {
	if isDebugMode {
		color.Enable()
	}
	level, err := zerolog.ParseLevel(levelstr)
	if err != nil {
		level = zerolog.InfoLevel
	}
	zerolog.TimeFieldFormat = time.DateTime
	zerolog.LevelFieldMarshalFunc = func(l zerolog.Level) string {
		return strings.ToUpper(l.String())
	}

	writers := []io.Writer{FileWriter}
	if isDebugMode {
		writers = append(writers, ConsoleWriter)
	}

	multi := zerolog.MultiLevelWriter(writers...)

	log.Logger = log.Output(multi).Level(level).With().Logger().Hook(TracingHook{})

	PanicLoggerInstance = zerolog.New(zerolog.MultiLevelWriter(multi)).Level(zerolog.ErrorLevel).With().Timestamp().Logger().Hook(TracingHook{})

}
