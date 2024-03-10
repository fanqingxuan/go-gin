package config

import (
	"go-gin/svc/redisx"
	"go-gin/svc/sqlx"
)

type Config struct {
	App   `yaml:"App"`
	Redis redisx.Config `yaml:"Redis"`
	Mysql sqlx.Config   `yaml:"Mysql"`
}
