package event

import (
	"go-gin/event/listener"
	"go-gin/internal/eventbus"
)

func Init() {
	eventbus.AddListener(SampleEventName, &listener.SampleAListener{}, &listener.SampleBListener{})
	eventbus.AddListener(DemoEventName, &listener.DemoAListener{})
}
