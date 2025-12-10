package event

import (
	"go-gin/internal/eventbus"
	"go-gin/model"
)

var DemoEventName eventbus.EventName = "event.demo"

func NewDemoEvent(u *model.User) *eventbus.Event {
	return eventbus.NewEvent(DemoEventName, u)
}
