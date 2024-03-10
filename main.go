package main

import (
	"flag"
	"fmt"
	"go-gin/config"
	"go-gin/handlers"
	"go-gin/middlewares"
	"go-gin/svc"
	"go-gin/utils/filex"

	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
)

var configFile = flag.String("f", "./.env.yaml", "the config file")

func main() {
	flag.Parse()
	setTimeDefault()
	var c config.Config
	filex.MustLoad(*configFile, &c)

	if c.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.New()
	server.HandleMethodNotAllowed = true

	serverCtx := svc.NewServiceContext(c)

	middlewares.RegisterGlobalMiddlewares(server)

	handlers.RegisterHandlers(server, serverCtx)

	fmt.Printf("Starting server at localhost%s...\n", c.App.Port)

	if err := server.Run(c.App.Port); err != nil {
		fmt.Printf("Start server error,err=%v", err)
	}
}

func setTimeDefault() {
	carbon.SetDefault(carbon.Default{
		Layout:       carbon.DateTimeLayout,
		Timezone:     carbon.PRC,
		WeekStartsAt: carbon.Sunday,
		Locale:       "zh-CN",
	})
}
