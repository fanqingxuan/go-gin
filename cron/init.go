package cron

import "go-gin/internal/scheduler"

func Init(cron *scheduler.MYCron) {

	cron.AddJob("@every 3s", &DBCheckJob{})
	cron.AddJob("@every 3s", &SampleJob{})
}
