package event

import "strings"

func eventName(event Event) string {
	name := strings.TrimSpace(event.Name())
	if name == "" {
		panic("event: the event name cannot be empty")
	}
	return name
}
