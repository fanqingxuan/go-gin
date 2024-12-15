package cron

import (
	"context"
	"go-gin/internal/components/db"
)

type DBCheckJob struct{}

func (j *DBCheckJob) Name() string {
	return "db check job"
}

func (j *DBCheckJob) Handle(ctx context.Context) error {
	if db.IsConnected() {
		return nil
	}
	return db.Connect()
}
