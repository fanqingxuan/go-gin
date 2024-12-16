package main

import (
	"flag"
	"go-gin/config"
	"go-gin/cron"
	"go-gin/event"
	"go-gin/internal/components/db"
	"go-gin/internal/components/logx"
	"go-gin/internal/components/redisx"
	"go-gin/internal/cronx"
)

var configFile = flag.String("f", "./.env", "the config file")

func main() {

	flag.Parse()

	config.Init(*configFile)
	config.InitGlobalVars()
	config.InitEnvironment()

	logx.InitConfig(config.GetLogConf())
	logx.Init()

	db.InitConfig(config.GetDbConf())
	db.Init()

	redisx.InitConfig(config.GetRedisConf())
	redisx.Init()

	event.Init()

	// 初始化第三方服务地址
	config.InitSvc()

	c := cronx.New()
	cron.Init(c)
	c.Run()

}
