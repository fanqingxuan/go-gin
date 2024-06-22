package config

import (
	"go-gin/internal/components/db"
	"go-gin/internal/components/logx"
	"go-gin/internal/components/redisx"
	"go-gin/internal/environment"
	filex "go-gin/internal/file"
	"go-gin/internal/ginx"
)

type App struct {
	Name     string           `yaml:"name"`
	Port     string           `yaml:"port"`
	Mode     environment.Mode `yaml:"mode"`
	TimeZone string           `yaml:"timezone"`
}

type Config struct {
	App   App           `yaml:"app"`
	Redis redisx.Config `yaml:"redis"`
	DB    db.Config     `yaml:"db"`
	Log   logx.Config   `yaml:"log"`
}

var instance Config

func Init(filename string) {
	err := filex.MustLoad(filename, &instance)
	if err != nil {
		panic(err)
	}
	environment.SetEnvMode(instance.App.Mode)
	environment.SetTimeZone(instance.App.TimeZone)
	ginx.InitConfig(ginx.Config{Port: instance.App.Port})

	logx.InitConfig(instance.Log)
	redisx.InitConfig(instance.Redis)
	db.InitConfig(instance.DB)
}

func GetAppConf() App {
	return instance.App
}
