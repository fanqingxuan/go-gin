package main

import (
	"flag"
	"go-gin/config"
	"go-gin/controllers"
	"go-gin/events"
	"go-gin/internal/components/db"
	"go-gin/internal/components/logx"
	"go-gin/internal/components/redisx"
	"go-gin/internal/task"
	_ "go-gin/internal/utils"

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

	InitConfig(Config{Port: config.GetAppConf().Port})
	engine := InitServer()
	InitValidators()
	controllers.Init(engine)
	StartServer(engine)

}
