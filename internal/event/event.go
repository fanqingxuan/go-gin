package event

import "context"

type Event struct {
	name    string
	payload any
}

func NewEvent(name string, payload any) *Event {
	return &Event{
		name:    name,
		payload: payload,
	}
}

func (e *Event) Name() string {
	return e.name
}

func (e *Event) Payload() any {
	return e.payload
}

func (e *Event) Fire(ctx context.Context) {
	Fire(ctx, e)
}

// FireAsync 异步执行事件监听
func (e *Event) FireAsync(ctx context.Context) {
	FireAsync(ctx, e)
}
