package httpc

import (
	"time"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	base *resty.Client
}

func (c *Client) SetBaseURL(url string) *Client {
	c.base.SetBaseURL(url)
	return c
}

func (c *Client) SetTimeout(timeout time.Duration) *Client {
	c.base.SetTimeout(timeout)
	return c
}

func (c *Client) NewRequest() *Request {
	return &Request{
		base: c.base.NewRequest(),
	}
}
