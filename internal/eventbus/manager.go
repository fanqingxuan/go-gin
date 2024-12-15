package eventbus

import (
	"context"
	"sync"
)

var mappings map[string][]Listener

var once sync.Once

// AddListener 为事件添加监听器
func AddListener(eventname string, listener ...Listener) {
	name := eventName(eventname)
	once.Do(func() {
		mappings = make(map[string][]Listener, 1024)
	})

	mappings[name] = append(mappings[name], listener...)
}

// Fire 同步执行事件监听,如果前一个返回error则停止执行
func Fire(ctx context.Context, event *Event) {
	name := eventName(event.Name())
	listeners := mappings[name]
	for _, listener := range listeners {
		if err := listener.Handle(ctx, event); err != nil {
			break
		}
	}
}

// FireAsync 异步执行事件监听
func FireAsync(ctx context.Context, event *Event) {
	name := eventName(event.Name())
	listeners := mappings[name]
	for _, listener := range listeners {
		go listener.Handle(ctx, event)
	}
}
