package httpc

import (
	"context"
	"fmt"
	"go-gin/internal/errorx"

	"github.com/go-resty/resty/v2"
)

type Request struct {
	base *resty.Request
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

func (r *Request) AddFormData(key string, data []string) *Request {
	for _, v := range data {
		r.base.FormData.Add(key, v)
	}
	return r
}

func (r *Request) SetBody(body interface{}) *Request {
	r.base.SetBody(body)
	return r
}

func (r *Request) SetResult(res interface{}) *Request {
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

func (r *Request) Exec() error {
	resp, err := r.base.Send()
	if err != nil {
		return errorx.New(errorx.ErrCodeThirdAPIConnectFailed, err.Error())
	}
	if resp.String() == "" {
		return errorx.New(errorx.ErrCodeThirdAPIContentNoContentFailed, "第三方接口返回数据为空")
	}
	if ret, ok := r.base.Result.(IBaseResponse); ok {
		if err := ret.Parse([]byte(resp.String())); err != nil {
			return errorx.New(errorx.ErrCodeThirdAPIContentParseFailed, fmt.Sprintf("第三方接口返回,解析响应内容失败,%w", err))
		}

		if !ret.Valid() {
			return errorx.New(errorx.ErrCodeThirdAPICallFormatFailed, "第三方接口返回数据格式错误")
		}
		if !ret.IsSuccess() {
			msg := ret.Msg()
			if msg == "" {
				msg = `第三方接口返回失败,但无返回提示消息`
			}
			return errorx.New(errorx.ErrCodeThirdAPIBusinessFailed, msg)
		}
		switch res := r.base.Result.(type) {
		case IRepsonseNonStardard:
			if err := res.ParseData([]byte(resp.String())); err != nil {
				return errorx.New(errorx.ErrCodeThirdAPIDataParseFailed, fmt.Sprintf("第三方接口已返回数据,解析数据失败,%w", err))
			}
		case IResponse:
			if err := res.ParseData(); err != nil {
				return errorx.New(errorx.ErrCodeThirdAPIDataParseFailed, fmt.Sprintf("第三方接口已返回数据,解析数据失败,%w", err))
			}
		}
	}
	return nil
}
