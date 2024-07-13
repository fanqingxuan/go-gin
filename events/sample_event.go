package events

import (
	"go-gin/internal/event"
)

type SampleEvent struct {
	name    string
	payload any
}

func CreateSampleEvent(user string) event.Event {
	return &SampleEvent{
		name:    "sample.event",
		payload: user,
	}
}

func (e *SampleEvent) Name() string {
	return e.name
}
func (e *SampleEvent) Payload() any {
	return e.payload
}
