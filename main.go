package main

import (
	"flag"
	"fmt"
	"go-gin/config"
	"go-gin/controllers"
	"go-gin/internal/ginx"
	"go-gin/internal/ginx/httpx"
	"go-gin/middlewares"
	"go-gin/pkg/db"
	"go-gin/pkg/logx"
	"go-gin/pkg/redisx"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

var configFile = flag.String("f", "./.env.yaml", "the config file")

func main() {
	flag.Parse()

	config.Init(*configFile)

	logx.Init(config.LogLevel(), config.IsDebugMode())

	db.Init()

	redisConfig := config.GetRedis()
	options := &redis.Options{
		Addr:     redisConfig.Addr,
		Username: redisConfig.Username,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.DB,       // use default DB
	}
	redisx.Init(options)

	httpx.DefaultSuccessCodeValue = 0

	engine := ginx.Init(config.IsDebugMode())
	middlewares.Init(engine)
	controllers.Init(engine)

	fmt.Printf("Starting server at localhost%s...\n", config.Port())
	if err := engine.Run(config.Port()); err != nil {
		fmt.Printf("Start server error,err=%v", err)
	}
}
