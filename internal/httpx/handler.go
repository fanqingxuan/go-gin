package httpx

import (
	"context"
	"errors"
	"go-gin/internal/component/logx"
	"io"

	"github.com/gin-gonic/gin/binding"
)

// LogicHandler 处理请求的逻辑接口
type LogicHandler[Req any, Resp any] interface {
	Handle(ctx context.Context, req Req) (Resp, error)
}

// Handle 自动根据 Content-Type 绑定参数并处理请求
func Handle[Req any, Resp any](c *Context, logicHandler LogicHandler[Req, Resp]) (Resp, error) {
	b := binding.Default(c.Request.Method, c.ContentType())
	return HandleWith(c, logicHandler, b)
}

// HandleJSON JSON 绑定参数并处理请求
func HandleJSON[Req any, Resp any](c *Context, logicHandler LogicHandler[Req, Resp]) (Resp, error) {
	return HandleWith(c, logicHandler, binding.JSON)
}

// HandleQuery Query 参数绑定并处理请求
func HandleQuery[Req any, Resp any](ctx *Context, logicHandler LogicHandler[Req, Resp]) (Resp, error) {
	return HandleWith(ctx, logicHandler, binding.Query)
}

// HandleHeader Header 绑定参数并处理请求
func HandleHeader[Req any, Resp any](ctx *Context, logicHandler LogicHandler[Req, Resp]) (Resp, error) {
	return HandleWith(ctx, logicHandler, binding.Header)
}

// HandleUri URI 参数绑定并处理请求
func HandleUri[Req any, Resp any](ctx *Context, logicHandler LogicHandler[Req, Resp]) (Resp, error) {
	var req Req
	var resp Resp
	if err := ctx.ShouldBindUri(&req); err != nil {
		if errors.Is(err, io.EOF) {
			logx.WithContext(ctx).Warn("ShouldBindUri异常", "io.EOF错误")
			return resp, err
		}
		logx.WithContext(ctx).Warn("ShouldBindUri异常", err)
		return resp, err
	}
	return logicHandler.Handle(ctx.Request.Context(), req)
}

// HandleWith 使用指定 Binding 绑定参数并处理请求
func HandleWith[Req any, Resp any](ctx *Context, logicHandler LogicHandler[Req, Resp], b binding.Binding) (Resp, error) {
	var req Req
	var resp Resp
	if err := ctx.ShouldBindWith(&req, b); err != nil {
		if errors.Is(err, io.EOF) {
			logx.WithContext(ctx).Warn("ShouldBind异常", "io.EOF错误")
			return resp, err
		}
		logx.WithContext(ctx).Warn("ShouldBind异常", err.Error())
		return resp, err
	}
	return logicHandler.Handle(ctx.Request.Context(), req)
}
