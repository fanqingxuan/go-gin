package event

import (
	"go-gin/internal/eventbus"
	"go-gin/model/entity"
)

var DemoEventName eventbus.EventName = "event.demo"

func NewDemoEvent(u *entity.User) *eventbus.Event {
	return eventbus.NewEvent(DemoEventName, u)
}
