package logx

import (
	"go-gin/internal/environment"
	"io"
	"strings"
	"time"

	"github.com/labstack/gommon/color"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	AccessLoggerInstance zerolog.Logger
	DBLoggerInstance     zerolog.Logger
	RestyLoggerInstance  zerolog.Logger
)

type Config struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"`
}

var conf Config

func InitConfig(c Config) {
	conf = c
}

func Init() {
	if environment.IsDebugMode() {
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

	fileWriter := zerolog.SyncWriter(&FileLevelWriter{
		Dirname:     conf.Path,
		FilePattern: time.DateOnly,
	})

	writers := []io.Writer{fileWriter}
	if environment.IsDebugMode() {
		writers = append(writers, &ConsoleLevelWriter{})
	}

	multi := zerolog.MultiLevelWriter(writers...)
	log.Logger = log.Output(multi).Level(level).With().Logger().Hook(TracingHook{})

	accessFileWriter := zerolog.SyncWriter(&FileLevelWriter{
		Dirname:     conf.Path + "access/",
		FilePattern: time.DateOnly,
	})
	AccessLoggerInstance = zerolog.New(accessFileWriter).Level(zerolog.InfoLevel).With().Timestamp().Logger().Hook(TracingHook{})

	restyFileWriter := zerolog.SyncWriter(&FileLevelWriter{
		Dirname:     conf.Path + "resty/",
		FilePattern: time.DateOnly,
	})
	RestyLoggerInstance = zerolog.New(restyFileWriter).Level(zerolog.InfoLevel).With().Timestamp().Logger().Hook(TracingHook{})
}
