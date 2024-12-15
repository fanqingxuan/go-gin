package main

import (
	"flag"
	"fmt"
	"go-gin/config"
	"go-gin/events"
	"go-gin/internal/components/db"
	"go-gin/internal/components/logx"
	"go-gin/internal/components/redisx"
	"go-gin/internal/environment"
	"go-gin/internal/httpx"
	"go-gin/internal/httpx/validators"
	"go-gin/internal/task"
	_ "go-gin/internal/utils"
	"go-gin/middleware"
	"go-gin/router"

	_ "github.com/go-sql-driver/mysql"
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
	task.Init(config.GetRedisConf())
	defer task.Close()
	events.Init()

	// 初始化第三方服务地址
	config.InitSvc()
	// 启动http服务
	startHttpServer(config.GetAppConf().Port)

}

func startHttpServer(port string) {
	if environment.IsDebugMode() {
		httpx.SetDebugMode()
	} else {
		httpx.SetReleaseMode()
	}
	engine := httpx.Default()
	validators.Init()
	middleware.Init(engine)
	router.Init(engine)
	fmt.Printf("Starting server at localhost%s...\n", port)
	if err := engine.Run(port); err != nil {
		fmt.Printf("Start server error,err=%v", err)
	}
}
