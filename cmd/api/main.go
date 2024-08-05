package main

import (
	"flag"
	"go-gin/config"
	"go-gin/controllers"
	"go-gin/events"
	"go-gin/internal/components/db"
	"go-gin/internal/components/logx"
	"go-gin/internal/components/redisx"
	"go-gin/internal/environment"
	"go-gin/internal/ginx"
	_ "go-gin/internal/utils"
	"go-gin/middlewares"
	"go-gin/rest/userc"

	_ "github.com/go-sql-driver/mysql"
)

var configFile = flag.String("f", "./.env", "the config file")

func main() {

	flag.Parse()

	config.Init(*configFile)
	config.InitGlobalVars()

	environment.SetEnvMode(config.GetAppConf().Mode)
	environment.SetTimeZone(config.GetAppConf().TimeZone)

	logx.InitConfig(config.GetLogConf())
	logx.Init()

	db.InitConfig(config.GetDbConf())
	db.Init()

	redisx.InitConfig(config.GetRedisConf())
	redisx.Init()

	events.Init()

	// 初始化第三方请求服务
	userc.InitUserSvc("http://localhost:8080/")

	ginx.InitConfig(ginx.Config{Port: config.GetAppConf().Port})
	engine := ginx.Init()
	middlewares.Init(engine)
	controllers.Init(engine)
	ginx.Start(engine)

}
