package event

import "context"

type Listener interface {
	Handle(context.Context, *Event) error
}
