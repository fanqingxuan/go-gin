package ginx

import (
	"fmt"
	"go-gin/internal/environment"
	"go-gin/internal/ginx/middlewares"
	"go-gin/internal/ginx/validators"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port string
}

var conf Config

func InitConfig(c Config) {
	conf = c
}

func Init() *gin.Engine {
	if environment.IsDebugMode() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.New()
	engine.HandleMethodNotAllowed = true
	engine.Use(middlewares.TraceId())
	engine.Use(middlewares.RequestLog())
	validators.Init()
	return engine
}

func Start(engine *gin.Engine) {
	fmt.Printf("Starting server at localhost%s...\n", conf.Port)
	if err := engine.Run(conf.Port); err != nil {
		fmt.Printf("Start server error,err=%v", err)
	}
}
