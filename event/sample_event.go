package event

import (
	"go-gin/internal/eventbus"
)

var SampleEventName = "event.sample"

func NewSampleEvent(user string) *eventbus.Event {
	return eventbus.NewEvent(SampleEventName, user)
}
