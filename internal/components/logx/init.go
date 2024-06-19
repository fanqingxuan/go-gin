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

	FileWriter io.Writer

	AccessFileWriter io.Writer
)

var (
	AccessLoggerInstance zerolog.Logger
	PanicLoggerInstance  zerolog.Logger
)

type Config struct {
	Level       string
	Path        string
	IsDebugMode bool
}

func Init(conf Config) {
	if conf.IsDebugMode {
		color.Enable()
	}
	level, err := zerolog.ParseLevel(conf.Level)
	if err != nil {
		level = zerolog.InfoLevel
	}
	zerolog.TimeFieldFormat = time.DateTime
	zerolog.LevelFieldMarshalFunc = func(l zerolog.Level) string {
		return strings.ToUpper(l.String())
	}

	FileWriter = zerolog.SyncWriter(&FileLevelWriter{
		Dirname:     conf.Path,
		FilePattern: time.DateOnly,
	})

	AccessFileWriter = zerolog.SyncWriter(&FileLevelWriter{
		Dirname:     conf.Path + "access/",
		FilePattern: time.DateOnly,
	})

	writers := []io.Writer{FileWriter}
	if conf.IsDebugMode {
		writers = append(writers, ConsoleWriter)
	}

	multi := zerolog.MultiLevelWriter(writers...)

	log.Logger = log.Output(multi).Level(level).With().Logger().Hook(TracingHook{})

	AccessLoggerInstance = zerolog.New(AccessFileWriter).Level(zerolog.InfoLevel).With().Timestamp().Logger().Hook(TracingHook{})
	PanicLoggerInstance = zerolog.New(zerolog.MultiLevelWriter(multi)).Level(zerolog.ErrorLevel).With().Timestamp().Logger().Hook(TracingHook{})

}
