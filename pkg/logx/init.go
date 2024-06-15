package logx

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	ConsoleWriter io.Writer = os.Stdout
)

var (
	AccessLoggerInstance zerolog.Logger = zerolog.New(ConsoleWriter).Level(zerolog.InfoLevel).With().Timestamp().Logger().Hook(TracingHook{})
	PanicLoggerInstance                 = zerolog.New(ConsoleWriter).Level(zerolog.ErrorLevel).With().Timestamp().Logger().Hook(TracingHook{})
)

func Init() {
	zerolog.TimeFieldFormat = time.DateTime
	zerolog.LevelFieldMarshalFunc = func(l zerolog.Level) string {
		return strings.ToUpper(l.String())
	}
	// zerolog.MessageFieldName = "m"

	multi := zerolog.MultiLevelWriter(ConsoleWriter)

	log.Logger = log.Output(multi).Level(zerolog.DebugLevel).With().Timestamp().Logger().Hook(TracingHook{})
}