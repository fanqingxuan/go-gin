package cron

import (
	"context"
	"go-gin/internal/component/db"
)

type DBCheckJob struct{}

func (j *DBCheckJob) Handle(ctx context.Context) error {
	if db.IsConnected() {
		return nil
	}
	return db.Connect()
}
