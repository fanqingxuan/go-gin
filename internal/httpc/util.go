package httpc

import (
	"context"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

type M map[string]string

var (
	defaultClient *Client
	once          sync.Once
)

func getDefaultClient() *Client {
	once.Do(func() {
		defaultClient = NewClient()
	})
	return defaultClient
}

func NewClient() *Client {
	client := &Client{
		base: resty.New(),
	}
	client.SetTimeout(3 * time.Minute)

	client.base.OnBeforeRequest(LogBeforeRequest)
	client.base.OnError(LogErrorHook)
	client.base.OnSuccess(LogSuccessHook)
	client.base.OnPanic(LogErrorHook)
	client.base.OnInvalid(LogErrorHook)
	return client
}

func GET(ctx context.Context, url string) *Request {
	return getDefaultClient().
		NewRequest().
		GET(url).
		SetContext(ctx)
}

func POST(ctx context.Context, url string) *Request {
	return getDefaultClient().
		NewRequest().
		POST(url).
		SetContext(ctx)
}
