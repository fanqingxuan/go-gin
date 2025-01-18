package cron

import (
	"context"
	"go-gin/internal/components/db"
)

type DBCheckJob struct{}

func (j *DBCheckJob) Handle(ctx context.Context) error {
	if db.IsConnected() {
		return nil
	}
	return db.Connect()
}
