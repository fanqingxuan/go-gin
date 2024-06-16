package config

import (
	"go-gin/internal/components/db"
	filex "go-gin/internal/file"
	"os"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

var instance Config

func Init(filename string) {
	err := filex.MustLoad(filename, &instance)
	if err != nil {
		panic(err)
	}
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

func GetRedisConf() *redis.Options {
	redisConfig := instance.Redis
	return &redis.Options{
		Addr:     redisConfig.Addr,
		Username: redisConfig.Username,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.DB,       // use default DB
	}
}

func GetDBConf() db.Config {
	return db.Config{
		DSN:          instance.DB.DSN,
		MaxOpenConns: instance.DB.MaxOpenConns,
		MaxIdleConns: instance.DB.MaxIdleConns,
	}
}

func LoadTimeZone() {
	err := os.Setenv("TZ", instance.TimeZone)
	if err != nil {
		panic("设置环境变量失败:" + err.Error())
	}
}
