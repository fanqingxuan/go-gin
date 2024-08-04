package request

import (
	"context"

	"github.com/go-resty/resty/v2"
)

type Request struct {
	base *resty.Request
}

func (r *Request) SetContentType(contentType string) *Request {
	r.base.ForceContentType(contentType)
	return r
}

func (r *Request) SetHeader(header, value string) *Request {
	r.base.SetHeader(header, value)
	return r
}

func (r *Request) SetHeaders(headers map[string]string) *Request {
	r.base.SetHeaders(headers)
	return r
}

func (r *Request) SetHeaderMultiValues(headers map[string][]string) *Request {
	r.base.SetHeaderMultiValues(headers)
	return r
}

func (r *Request) SetQueryString(query string) *Request {
	r.base.SetQueryString(query)
	return r
}

func (r *Request) SetQueryParam(param, value string) *Request {
	r.base.SetQueryParam(param, value)
	return r
}

func (r *Request) SetQueryParams(params map[string]string) *Request {
	r.base.SetQueryParams(params)
	return r
}

func (r *Request) SetFormData(data map[string]string) *Request {
	r.base.SetFormData(data)
	return r
}

func (r *Request) SetBody(body interface{}) *Request {
	r.base.SetBody(body)
	return r
}

func (r *Request) ParseResult(res interface{}) *Request {
	r.base.SetResult(res)
	return r
}

func (r *Request) SetContext(ctx context.Context) *Request {
	r.base.SetContext(ctx)
	return r
}

func (r *Request) GET(url string) *Request {
	r.base.Method = resty.MethodGet
	r.base.URL = url
	return r
}

func (r *Request) POST(url string) *Request {
	r.base.Method = resty.MethodPost
	r.base.URL = url
	return r
}

func (r *Request) Send() (*resty.Response, error) {
	return r.base.Send()
}
