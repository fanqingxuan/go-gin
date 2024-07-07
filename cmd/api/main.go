package main

import (
	"flag"
	"go-gin/config"
	"go-gin/controllers"
	"go-gin/internal/components/db"
	"go-gin/internal/components/logx"
	"go-gin/internal/components/redisx"
	"go-gin/internal/environment"
	"go-gin/internal/ginx"
	_ "go-gin/internal/utils"
	"go-gin/middlewares"

	_ "github.com/go-sql-driver/mysql"
)

var configFile = flag.String("f", "./.env", "the config file")

func main() {

	flag.Parse()

	config.Init(*configFile)
	config.InitGlobalVars()

	environment.SetEnvMode(config.Instance.App.Mode)
	environment.SetTimeZone(config.Instance.App.TimeZone)

	logx.InitConfig(config.Instance.Log)
	logx.Init()

	db.InitConfig(config.Instance.DB)
	db.Init()

	redisx.InitConfig(config.Instance.Redis)
	redisx.Init()

	ginx.InitConfig(ginx.Config{Port: config.Instance.App.Port})
	engine := ginx.Init()
	middlewares.Init(engine)
	controllers.Init(engine)
	ginx.Start(engine)

}
