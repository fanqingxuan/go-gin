package job

import "go-gin/internal/cron"

func Init(cron *cron.MYCron) {

	cron.AddJob("@every 3s", &SampleJob{})
}
