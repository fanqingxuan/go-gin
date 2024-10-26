package events

import (
	"go-gin/internal/event"
)

var SampleEventName = "event.sample"

func NewSampleEvent(user string) *event.Event {
	return event.NewEvent(SampleEventName, user)
}
