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

	"github.com/golang-module/carbon/v2"
)

type App struct {
	Name       string           `yaml:"name"`
	Port       string           `yaml:"port"`
	Mode       environment.Mode `yaml:"mode"`
	TimeZone   string           `yaml:"timezone"`
	TimeFormat string           `yaml:"timeformat"`
}

type SvcConfig struct {
	UserSvcUrl  string `yaml:"user_url"`
	LoginSvcUrl string `yaml:"login_url"`
}

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

func InitGlobalVars() {
	// httpx.DefaultSuccessCodeValue = 0
	httpx.DefaultSuccessMessageValue = "成功"

	traceid.TraceIdFieldName = "requestId"
}

func InitEnvironment() {

	carbon.SetDefault(carbon.Default{
		Layout:   instance.App.TimeFormat,
		Timezone: instance.App.TimeZone,
	})
	environment.SetEnvMode(instance.App.Mode)
	environment.SetTimeZone(instance.App.TimeZone)
}

func GetAppConf() App {
	return instance.App
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

func GetSvcConfig() SvcConfig {
	return instance.Svc
}
