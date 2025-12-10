package eventbus

import (
	"context"
	"sync"
)

var (
	mappings map[EventName][]Listener
	mu       sync.RWMutex
	once     sync.Once
)

// AddListener 为事件添加监听器
func AddListener(eventname EventName, listener ...Listener) {
	name := validateEventName(eventname)
	once.Do(func() {
		mappings = make(map[EventName][]Listener, 1024)
	})

	mu.Lock()
	defer mu.Unlock()
	mappings[name] = append(mappings[name], listener...)
}

// Fire 同步执行事件监听,如果前一个返回error则停止执行
func Fire(ctx context.Context, event *Event) {
	name := validateEventName(event.Name())
	mu.RLock()
	listeners := mappings[name]
	mu.RUnlock()

	for _, listener := range listeners {
		if err := listener.Handle(ctx, event); err != nil {
			break
		}
	}
}

// FireIf 同步执行事件监听,如果第一个参数是true则运行事件
func FireIf(ctx context.Context, condition bool, event *Event) {
	if condition {
		Fire(ctx, event)
	}
}

// FireAsync 异步执行事件监听
func FireAsync(ctx context.Context, event *Event) {
	name := validateEventName(event.Name())
	mu.RLock()
	listeners := mappings[name]
	mu.RUnlock()

	for _, listener := range listeners {
		l := listener // 避免闭包问题
		go l.Handle(ctx, event)
	}
}

// FireAsyncIf 异步执行事件监听,如果第一个参数是true则运行事件
func FireAsyncIf(ctx context.Context, condition bool, event *Event) {
	if condition {
		FireAsync(ctx, event)
	}
}
