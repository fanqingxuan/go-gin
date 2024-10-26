package events

import (
	"go-gin/internal/event"
	"go-gin/listeners"
)

func Init() {
	event.AddListener(SampleEventName, &listeners.SampleAListener{}, &listeners.SampleBListener{})
	event.AddListener(DemoEventName, &listeners.DemoAListener{})
}
