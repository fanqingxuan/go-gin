package main

import (
	"flag"
	"go-gin/config"
	"go-gin/cron"
	"go-gin/event"
	"go-gin/internal/component/db"
	"go-gin/internal/component/logx"
	"go-gin/internal/component/redisx"
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

	// 定时任务
	cronx.New()
	cron.Init()
	cronx.Run()

}
