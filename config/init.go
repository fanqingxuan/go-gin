package config

import (
	"go-gin/internal/components/db"
	"go-gin/internal/components/logx"
	filex "go-gin/internal/file"
	"os"
	"path/filepath"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	App   `yaml:"app"`
	Redis `yaml:"redis"`
	DB    `yaml:"db"`
	Log   `yaml:"log"`
}

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

func GetLogConf() logx.Config {
	return logx.Config{
		Level:       instance.Log.Level,
		Path:        filepath.ToSlash(instance.Log.Path) + "/",
		IsDebugMode: IsDebugMode(),
	}
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
		LogLevel:     instance.Log.Level,
	}
}

func LoadTimeZone() {
	err := os.Setenv("TZ", instance.TimeZone)
	if err != nil {
		panic("设置环境变量失败:" + err.Error())
	}
}
