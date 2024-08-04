package request

import (
	"go-gin/internal/components/logx"

	"github.com/go-resty/resty/v2"
)

type Hook interface {
	Before(*resty.Client, *resty.Request) error
	After(*resty.Client, *resty.Response) error
}

type ErrorHook interface {
	Handle(*resty.Request, error)
}

type LogHook struct {
}

func (h *LogHook) Before(c *resty.Client, r *resty.Request) error {
	logx.RestyLoggerInstance.Info().Ctx(r.Context()).
		Str("keywords", "request").
		Str("url", r.URL).
		Str("method", r.Method).
		Any("header", r.Header).
		Str("query", r.QueryParam.Encode()).
		Any("post", r.FormData).
		Any("body", r.Body).
		Send()
	return nil
}

func (h *LogHook) After(c *resty.Client, r *resty.Response) error {
	logx.RestyLoggerInstance.Info().Ctx(r.Request.Context()).
		Str("keywords", "response").
		Int("code", r.StatusCode()).
		Str("status", r.Status()).
		Any("proto", r.Proto()).
		Any("header", r.Header()).
		Any("error", r.Error()).
		Str("body", r.String()).
		Any("traceInfo", r.Request.TraceInfo()).
		Send()
	return nil
}

func (h *LogHook) Handle(r *resty.Request, err error) {
	logx.RestyLoggerInstance.Info().Ctx(r.Context()).
		Str("keywords", "reqeust or response error").
		Str("msg", err.Error()).
		Send()
}
