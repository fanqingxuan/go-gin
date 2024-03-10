package config

import "go-gin/svc/redisx"

type Config struct {
	App   `yaml:"App"`
	Redis redisx.Config `yaml:"Redis"`
	Mysql `yaml:"Mysql"`
}
