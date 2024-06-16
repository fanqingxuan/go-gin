package main

import (
	"flag"
	"go-gin/config"
	"go-gin/internal/components/db"
	"go-gin/internal/components/logx"
	"go-gin/internal/components/redisx"
	"go-gin/internal/cron"
	"go-gin/internal/ginx/httpx"
	"go-gin/job"
)

var configFile = flag.String("f", "./.env.yaml", "the config file")

func main() {
	c := cron.New()
	flag.Parse()

	config.Init(*configFile)

	config.LoadTimeZone()

	logx.Init(config.LogLevel(), config.IsDebugMode())

	db.Init(config.GetDBConf())

	redisx.Init(config.GetRedisConf())

	httpx.DefaultSuccessCodeValue = 0

	job.Init(c)
	// c.AddFunc("%every 1s", func() { fmt.Println("Every hour on the half hour") })
	// c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	// c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })

	c.Run()

}
