package main

import (
	"flag"
	"fmt"
	"go-gin/config"
	"go-gin/internal/component/db"
	"go-gin/internal/component/logx"
	"go-gin/internal/component/redisx"
	"go-gin/internal/queue"
	"go-gin/task"
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
	queue.InitServer(config.GetRedisConf())
	task.Init()
	if err := queue.Start(); err != nil {
		fmt.Printf("could not run server: %v", err)
	}
}
