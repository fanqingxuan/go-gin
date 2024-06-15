package logx

import (
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
	AccessLoggerInstance zerolog.Logger = zerolog.New(zerolog.MultiLevelWriter(ConsoleWriter, AcessFileWriter)).Level(zerolog.InfoLevel).With().Timestamp().Logger().Hook(TracingHook{})
	PanicLoggerInstance                 = zerolog.New(zerolog.MultiLevelWriter(ConsoleWriter, FileWriter)).Level(zerolog.ErrorLevel).With().Timestamp().Logger().Hook(TracingHook{})
)

func Init(level zerolog.Level, isDebugMode bool) {
	if isDebugMode {
		color.Enable()
	}
	zerolog.TimeFieldFormat = time.DateTime
	zerolog.LevelFieldMarshalFunc = func(l zerolog.Level) string {
		return strings.ToUpper(l.String())
	}

	multi := zerolog.MultiLevelWriter(ConsoleWriter, FileWriter)

	log.Logger = log.Output(multi).Level(level).With().Logger().Hook(TracingHook{})
}
