package main

import (
	"flag"
	"go-gin/config"
	"go-gin/internal/components/db"
	"go-gin/internal/components/logx"
	"go-gin/internal/components/redisx"
	"go-gin/internal/cron"
	"go-gin/internal/environment"
	"go-gin/jobs"
)

var configFile = flag.String("f", "./.env", "the config file")

func main() {

	flag.Parse()

	config.Init(*configFile)
	config.InitGlobalVars()

	environment.SetEnvMode(config.Instance.App.Mode)
	environment.SetTimeZone(config.Instance.App.TimeZone)

	logx.InitConfig(config.Instance.Log)
	logx.Init()

	db.InitConfig(config.Instance.DB)
	db.Init()

	redisx.InitConfig(config.Instance.Redis)
	redisx.Init()

	c := cron.New()
	jobs.Init(c)
	c.Run()

}
