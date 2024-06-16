package job

import (
	"context"
	"go-gin/internal/components/db"
	"go-gin/models"
)

type SampleJob struct{}

func (j *SampleJob) Name() string {
	return "sample job"
}

func (j *SampleJob) Handle(ctx context.Context) error {

	var u models.User
	db.WithContext(ctx).Find(&u)

	return nil
}
