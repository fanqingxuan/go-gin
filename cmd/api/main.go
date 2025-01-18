package main

import (
	"flag"
	"fmt"
	"go-gin/config"
	"go-gin/event"
	"go-gin/internal/components/db"
	"go-gin/internal/components/logx"
	"go-gin/internal/components/redisx"
	"go-gin/internal/environment"
	"go-gin/internal/errorx"
	"go-gin/internal/httpx"
	"go-gin/internal/httpx/validators"
	"go-gin/internal/queue"
	_ "go-gin/internal/util"
	"go-gin/middleware"
	"go-gin/router"
	"go-gin/util"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hibiken/asynqmon"
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

	queue.Init(config.GetRedisConf())
	defer queue.Close()

	event.Init()

	// 初始化第三方服务地址
	config.InitSvc()

	// 初始化http服务
	engine := initHttpServer()

	// 挂载监控
	mountMonitor(engine)

	// 启动http服务
	port := config.GetAppConf().Port
	fmt.Printf("Starting server at localhost%s...\n", port)
	if err := engine.Run(port); err != nil {
		fmt.Printf("Start server error,err=%v", err)
	}
}

// 初始化http服务
func initHttpServer() *httpx.Engine {
	if environment.IsDebugMode() {
		httpx.SetDebugMode()
	} else {
		httpx.SetReleaseMode()
	}
	engine := httpx.Default()
	validators.Init()
	middleware.Init(engine)
	router.Init(engine)
	return engine
}

// 挂载监控
func mountMonitor(engine *httpx.Engine) {
	// 挂载队列监控web ui
	mon := asynqmon.New(asynqmon.Options{
		RootPath:     "/monitor/queue",
		RedisConnOpt: queue.RedisClientOpt(config.GetRedisConf()),
	})

	r := engine.Engine.Group("/monitor")
	r.Use(func(ctx *gin.Context) {
		clientIp := ctx.ClientIP()
		// 如果客户端ip不在白名单内，直接返回403
		if !util.InArray(clientIp, config.GetMonitorConf().WhiteIpList) {
			httpx.Error(httpx.NewContext(ctx), errorx.ErrorForbidden)
			ctx.Abort()
			return
		}
		ctx.Next()
	})
	r.GET("/queue/*any", gin.WrapH(mon))
}
