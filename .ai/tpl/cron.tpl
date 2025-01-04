package cron

import (
	"context"
)

type XxxJob struct{}

func (j *XxxJob) Name() string {
	return "xxx job"
}

func (j *XxxJob) Handle(ctx context.Context) error {
	return nil
}
