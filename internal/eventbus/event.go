package eventbus

import "context"

type EventName string

type Event struct {
	name    EventName
	payload any
}

func NewEvent(name EventName, payload any) *Event {
	return &Event{
		name:    name,
		payload: payload,
	}
}

func (e *Event) Name() EventName {
	return e.name
}

func (e *Event) Payload() any {
	return e.payload
}

func (e *Event) Fire(ctx context.Context) {
	Fire(ctx, e)
}

func (e *Event) FireIf(ctx context.Context, condition bool) {
	FireIf(ctx, condition, e)
}

// FireAsync 异步执行事件监听
func (e *Event) FireAsync(ctx context.Context) {
	FireAsync(ctx, e)
}

// FireAsyncIf 异步执行事件监听,如果第一个参数是true则运行事件
func (e *Event) FireAsyncIf(ctx context.Context, condition bool) {
	FireAsyncIf(ctx, condition, e)
}
