package main

import (
	"flag"
	"go-gin/config"
	"go-gin/controllers"
	"go-gin/events"
	"go-gin/internal/components/db"
	"go-gin/internal/components/logx"
	"go-gin/internal/components/redisx"
	"go-gin/internal/ginx"
	_ "go-gin/internal/utils"
	"go-gin/middlewares"
	"go-gin/rest/login"
	"go-gin/rest/mylogin"
	"go-gin/rest/user"

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

	events.Init()

	// 初始化第三方服务地址
	InitSvc()

	ginx.InitConfig(ginx.Config{Port: config.GetAppConf().Port})
	engine := ginx.Init()
	middlewares.Init(engine)
	controllers.Init(engine)
	ginx.Start(engine)

}

func InitSvc() {
	svcConfig := config.GetSvcConfig()
	// 初始化第三方请求服务
	user.Init(svcConfig.UserSvcUrl)
	login.Init(svcConfig.LoginSvcUrl)
	mylogin.Init(svcConfig.LoginSvcUrl)
}
