package cron

import (
	"context"
)

type SampleJob struct{}

func (j *SampleJob) Handle(ctx context.Context) error {
	return nil
}
