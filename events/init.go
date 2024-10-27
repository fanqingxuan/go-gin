package events

import (
	"go-gin/events/listeners"
	"go-gin/internal/event"
)

func Init() {
	event.AddListener(SampleEventName, &listeners.SampleAListener{}, &listeners.SampleBListener{})
	event.AddListener(DemoEventName, &listeners.DemoAListener{})
}
