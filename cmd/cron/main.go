package main

import (
	"flag"
	"go-gin/config"
	"go-gin/internal/components/db"
	"go-gin/internal/components/logx"
	"go-gin/internal/components/redisx"
	"go-gin/internal/cron"
	"go-gin/jobs"
)

var configFile = flag.String("f", "./.env", "the config file")

func main() {
	c := cron.New()
	flag.Parse()

	config.Init(*configFile)
	config.InitGlobalVars()

	logx.Init()
	db.Init()
	redisx.Init()

	jobs.Init(c)

	c.Run()

}
