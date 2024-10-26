package main

import (
	"flag"
	"fmt"
	"go-gin/config"
	"go-gin/internal/components/db"
	"go-gin/internal/components/logx"
	"go-gin/internal/components/redisx"
	"go-gin/internal/task"
	"go-gin/tasks"
)

var configFile = flag.String("f", "./.env", "the config file")

func main() {

	flag.Parse()

	config.Init(*configFile)
	config.InitEnvironment()

	logx.InitConfig(config.GetLogConf())
	logx.Init()

	db.InitConfig(config.GetDbConf())
	db.Init()

	redisx.InitConfig(config.GetRedisConf())
	redisx.Init()
	task.InitServer(config.GetRedisConf())
	tasks.Init()
	if err := task.Start(); err != nil {
		fmt.Printf("could not run server: %v", err)
	}
}
