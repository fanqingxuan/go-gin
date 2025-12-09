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
	CronLoggerInstance   zerolog.Logger
	QueueLoggerInstance  zerolog.Logger
)

type Config struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"`
}

var conf Config

func InitConfig(c Config) {
	if !strings.HasSuffix(c.Path, "/") {
		c.Path = c.Path + "/"
	}
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

	log.Logger = initDefaultInstance(level)
	AccessLoggerInstance = initLoggerInstance("access")
	RestyLoggerInstance = initLoggerInstance("httpc")
	CronLoggerInstance = initLoggerInstance("access_cron")
	QueueLoggerInstance = initLoggerInstance("access_queue")
}

func initDefaultInstance(l zerolog.Level) zerolog.Logger {
	fileWriter := zerolog.SyncWriter(&FileLevelWriter{
		Dirname:     conf.Path,
		FilePattern: time.DateOnly,
	})

	writers := []io.Writer{fileWriter}
	if environment.IsDebugMode() {
		writers = append(writers, &ConsoleLevelWriter{})
	}
	multi := zerolog.MultiLevelWriter(writers...)
	return zerolog.New(multi).Level(l).With().Timestamp().Logger().Hook(TracingHook{})
}

func initLoggerInstance(path string) zerolog.Logger {
	queueFileWriter := zerolog.SyncWriter(&FileLevelWriter{
		Dirname:     conf.Path + strings.Trim(path, "/") + "/",
		FilePattern: time.DateOnly,
	})
	writers := []io.Writer{queueFileWriter}
	if environment.IsDebugMode() {
		writers = append(writers, &ConsoleLevelWriter{})
	}
	return zerolog.New(zerolog.MultiLevelWriter(writers...)).Level(zerolog.InfoLevel).With().Timestamp().Logger().Hook(TracingHook{})
}
