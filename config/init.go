package config

import (
	"go-gin/internal/components/db"
	"go-gin/internal/components/logx"
	"go-gin/internal/components/redisx"
	filex "go-gin/internal/file"
	"sync"
)

type Config struct {
	App   App           `yaml:"app"`
	Redis redisx.Config `yaml:"redis"`
	DB    db.Config     `yaml:"db"`
	Log   logx.Config   `yaml:"log"`
	Svc   SvcConfig     `yaml:"svc"`
}

var instance *Config
var once sync.Once

func Init(filename string) {
	once.Do(func() {
		err := filex.MustLoad(filename, &instance)
		if err != nil {
			panic(err)
		}
	})
}

func GetRedisConf() redisx.Config {
	return instance.Redis
}

func GetLogConf() logx.Config {
	return instance.Log
}

func GetDbConf() db.Config {
	return instance.DB
}
