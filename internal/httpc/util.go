package httpc

import (
	"context"
	"time"
)

type M map[string]string

func New() *Request {
	client := NewClient()

	client.SetTimeout(3 * time.Minute)

	client.AddErrorHook(&LogHook{})
	client.AddHook(&LogHook{})

	return client.NewRequest()
}

func GET(ctx context.Context, url string) *Request {
	return New().GET(url).SetContext(ctx)
}

func POST(ctx context.Context, url string) *Request {
	return New().POST(url).SetContext(ctx)
}

func PostBody(ctx context.Context, url string) *Request {
	return New().POST(url).SetContext(ctx)
}
