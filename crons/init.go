package crons

import "go-gin/internal/cron"

func Init(cron *cron.MYCron) {

	cron.AddJob("@every 3s", &DBCheckJob{})
	cron.AddJob("@every 3s", &SampleJob{})
}
