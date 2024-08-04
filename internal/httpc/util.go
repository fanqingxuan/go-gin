package httpc

import (
	"context"
	"time"

	"github.com/go-resty/resty/v2"
)

type M map[string]string

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
	return NewClient().
		NewRequest().
		GET(url).
		SetContext(ctx)
}

func POST(ctx context.Context, url string) *Request {
	return NewClient().
		NewRequest().
		POST(url).
		SetContext(ctx)
}
