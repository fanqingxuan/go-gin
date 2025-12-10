package eventbus

import "strings"

func validateEventName(s EventName) EventName {
	name := strings.TrimSpace(string(s))
	if name == "" {
		panic("event: the event name cannot be empty")
	}
	return EventName(name)
}
