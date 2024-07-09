package events

import (
	"context"
	"go-gin/internal/components/db"
	"go-gin/models"
)

type sampleEvent struct {
}

func CreateSampleEvent() Event {
	return &sampleEvent{}
}

func (e *sampleEvent) Handle(ctx context.Context) error {
	var u models.User
	if err := db.WithContext(ctx).Find(&u).Error; err != nil {
		return err
	}
	return nil
}
