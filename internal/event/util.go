package event

import "strings"

func eventName(s string) string {
	name := strings.TrimSpace(s)
	if name == "" {
		panic("event: the event name cannot be empty")
	}
	return name
}
