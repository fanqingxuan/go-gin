package events

import (
	"go-gin/internal/event"
	"go-gin/listeners"
)

func Init() {
	event.AddListener(CreateSampleEvent(""), &listeners.SampleAListener{}, &listeners.SampleBListener{})
}
