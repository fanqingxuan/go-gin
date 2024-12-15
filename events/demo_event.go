package events

import (
	"go-gin/internal/event"
	"go-gin/model"
)

var DemoEventName = "event.demo"

func NewDemoEvent(u *model.User) *event.Event {
	return event.NewEvent(DemoEventName, u)
}
