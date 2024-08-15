package config

import "go-gin/internal/environment"

type App struct {
	Name       string           `yaml:"name"`
	Port       string           `yaml:"port"`
	Mode       environment.Mode `yaml:"mode"`
	TimeZone   string           `yaml:"timezone"`
	TimeFormat string           `yaml:"timeformat"`
}

func GetAppConf() App {
	return instance.App
}
