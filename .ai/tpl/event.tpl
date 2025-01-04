package event

import (
	"go-gin/internal/eventbus"
)

var XxxEventName = "event.xxx"

func NewXxxEvent(payload 参数类型) *eventbus.Event {
	return eventbus.NewEvent(XxxEventName, payload)
}
