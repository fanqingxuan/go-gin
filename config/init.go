package config

import (
	filex "go-gin/internal/file"
	"os"

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

func GetDB() DB {
	return instance.DB
}

func LoadTimeZone() {
	err := os.Setenv("TZ", instance.TimeZone)
	if err != nil {
		panic("设置环境变量失败:" + err.Error())
	}
}
