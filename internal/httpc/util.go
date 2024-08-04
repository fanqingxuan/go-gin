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
	client.AddErrorHook(&LogHook{})
	client.AddHook(&LogHook{})
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
