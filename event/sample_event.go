package event

import (
	"go-gin/internal/eventbus"
)

var SampleEventName eventbus.EventName = "event.sample"

func NewSampleEvent(user string) *eventbus.Event {
	return eventbus.NewEvent(SampleEventName, user)
}
