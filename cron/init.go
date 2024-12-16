package cron

import (
	"go-gin/internal/cronx"
)

func Init(cron *cronx.MYCron) {

	cron.AddJob("@every 3s", &DBCheckJob{})
	cron.AddJob("@every 3s", &SampleJob{})
}
