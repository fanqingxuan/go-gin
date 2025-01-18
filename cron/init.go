package cron

import "go-gin/internal/cronx"

func Init() {

	cronx.AddJob("@every 3s", &DBCheckJob{})
	// cronx.AddJob("@every 3s", &SampleJob{})
	cronx.Schedule(&SampleJob{}).EveryMinute()
}
