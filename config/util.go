package config

import (
	"go-gin/internal/components/db"
	"go-gin/internal/components/logx"
	"go-gin/internal/components/redisx"
)

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
