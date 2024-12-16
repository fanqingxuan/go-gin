package main

import (
	"flag"
	"fmt"
	"go-gin/config"
	"go-gin/internal/components/db"
	"go-gin/internal/components/logx"
	"go-gin/internal/components/redisx"
	"go-gin/internal/taskx"
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
	taskx.InitServer(config.GetRedisConf())
	task.Init()
	if err := taskx.Start(); err != nil {
		fmt.Printf("could not run server: %v", err)
	}
}
