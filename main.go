package main

import (
	"flag"
	"fmt"
	"go-gin/config"
	"go-gin/controllers"
	filex "go-gin/internal/file"
	"go-gin/internal/ginx/httpx"
	"go-gin/internal/ginx/validators"
	"go-gin/middlewares"
	"go-gin/pkg/db"
	"go-gin/pkg/logx"
	"go-gin/pkg/redis"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var configFile = flag.String("f", "./.env.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	filex.MustLoad(*configFile, &c)
	logx.Init()
	if c.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.New()
	server.HandleMethodNotAllowed = true

	validators.Init()

	middlewares.Init(server)

	db.Init()

	redis.Init()

	httpx.DefaultSuccessCodeValue = 0

	controllers.Init(server)

	fmt.Printf("Starting server at localhost%s...\n", c.App.Port)

	if err := server.Run(c.App.Port); err != nil {
		fmt.Printf("Start server error,err=%v", err)
	}
}
