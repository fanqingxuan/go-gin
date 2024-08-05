package httpc

import "github.com/go-resty/resty/v2"

type IResponse interface {
	ParseResponse(res interface{}, resp *resty.Response) error
}
