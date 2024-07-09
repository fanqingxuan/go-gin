package events

import "context"

type Event interface {
	Handle(ctx context.Context) error
}

func Fire(ctx context.Context, e Event) error {
	return e.Handle(ctx)
}
