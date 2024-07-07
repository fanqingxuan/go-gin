package config

import (
	"go-gin/internal/components/db"
	"go-gin/internal/components/logx"
	"go-gin/internal/components/redisx"
	"go-gin/internal/environment"
	filex "go-gin/internal/file"
	"go-gin/internal/ginx/httpx"
	"go-gin/internal/traceid"
	"sync"
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

var Instance *Config
var once sync.Once

func Init(filename string) {
	once.Do(func() {
		err := filex.MustLoad(filename, &Instance)
		if err != nil {
			panic(err)
		}
	})
}

func InitGlobalVars() {
	httpx.DefaultSuccessCodeValue = 0
	httpx.DefaultSuccessMessageValue = "成功"

	traceid.TraceIdFieldName = "requestId"
}
