package events

import (
	"go-gin/internal/event"
	"go-gin/models"
)

var DemoEventName = "event.demo"

func NewDemoEvent(u *models.User) *event.Event {
	return event.NewEvent(DemoEventName, u)
}
