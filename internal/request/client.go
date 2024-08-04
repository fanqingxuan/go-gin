package request

import (
	"time"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	base *resty.Client
}

func NewClient() *Client {
	return &Client{
		base: resty.New(),
	}
}

func (c *Client) SetTimeout(timeout time.Duration) *Client {
	c.base.SetTimeout(timeout)
	return c
}

func (c *Client) AddHook(h Hook) *Client {
	c.base.OnBeforeRequest(h.Before)
	c.base.OnAfterResponse(h.After)
	return c
}
func (c *Client) AddErrorHook(h ErrorHook) *Client {
	c.base.OnError(h.Handle)
	return c
}

func (c *Client) NewRequest() *Request {
	return &Request{
		base: c.base.NewRequest(),
	}
}
