package crons

import (
	"context"
)

type SampleJob struct{}

func (j *SampleJob) Name() string {
	return "sample job"
}

func (j *SampleJob) Handle(ctx context.Context) error {
	return nil
}
