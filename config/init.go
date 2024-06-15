package config

import (
	filex "go-gin/internal/file"

	"github.com/rs/zerolog"
)

var instance Config

func Init(filename string) {
	filex.MustLoad(filename, &instance)
}

func IsDebugMode() bool {
	return instance.App.Debug
}

func Port() string {
	return instance.App.Port
}

func LogLevel() zerolog.Level {
	l, err := zerolog.ParseLevel(instance.Log.Level)
	if err != nil {
		panic(err)
	}
	return l
}

func GetRedis() Redis {
	return instance.Redis
}
