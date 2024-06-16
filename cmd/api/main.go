package main

import (
	"flag"
	"fmt"
	"go-gin/config"
	"go-gin/controllers"
	"go-gin/internal/components/db"
	"go-gin/internal/components/logx"
	"go-gin/internal/components/redisx"
	"go-gin/internal/ginx"
	"go-gin/internal/ginx/httpx"
	_ "go-gin/internal/utils"
	"go-gin/middlewares"

	_ "github.com/go-sql-driver/mysql"
)

var configFile = flag.String("f", "./.env", "the config file")

func main() {

	flag.Parse()

	config.Init(*configFile)

	config.LoadTimeZone()

	logx.Init(config.LogLevel(), config.IsDebugMode())

	db.Init(config.GetDBConf())

	redisx.Init(config.GetRedisConf())

	httpx.DefaultSuccessCodeValue = 0

	engine := ginx.Init(config.IsDebugMode())
	middlewares.Init(engine)
	controllers.Init(engine)

	fmt.Printf("Starting server at localhost%s...\n", config.Port())
	if err := engine.Run(config.Port()); err != nil {
		fmt.Printf("Start server error,err=%v", err)
	}
}
