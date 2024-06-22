package main

import (
	"flag"
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
	logx.Init()
	db.Init()
	redisx.Init()

	httpx.DefaultSuccessCodeValue = 0
	httpx.DefaultSuccessMessageValue = "成功"

	engine := ginx.Init()
	middlewares.Init(engine)
	controllers.Init(engine)

	ginx.Start(engine)

}
