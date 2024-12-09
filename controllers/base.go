package controllers

import (
	"context"
	"errors"
	"go-gin/internal/components/logx"
	"go-gin/internal/httpx"
	"io"

	"github.com/gin-gonic/gin/binding"
)

// LogicHandler 处理请求的逻辑接口
type LogicHandler[Req any, Resp any] interface {
	Handle(ctx context.Context, req Req) (Resp, error)
}

// ShouldBindHandle 处理请求
func ShouldBindHandle[Req any, Resp any](c *httpx.Context, logicHandler LogicHandler[Req, Resp]) (Resp, error) {
	b := binding.Default(c.Request.Method, c.ContentType())
	return ShouldBindWithHandle(c, logicHandler, b)
}

func ShouldBindJSONHandle[Req any, Resp any](c *httpx.Context, logicHandler LogicHandler[Req, Resp]) (Resp, error) {
	return ShouldBindWithHandle(c, logicHandler, binding.JSON)
}

func ShouldBindQueryHandle[Req any, Resp any](ctx *httpx.Context, logicHandler LogicHandler[Req, Resp]) (Resp, error) {
	return ShouldBindWithHandle(ctx, logicHandler, binding.Query)
}

func ShouldBindHeaderHandle[Req any, Resp any](ctx *httpx.Context, logicHandler LogicHandler[Req, Resp]) (Resp, error) {
	return ShouldBindWithHandle(ctx, logicHandler, binding.Header)
}

func ShouldBindUriHandle[Req any, Resp any](ctx *httpx.Context, logicHandler LogicHandler[Req, Resp]) (Resp, error) {
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
	return logicHandler.Handle(ctx, req)
}

func ShouldBindWithHandle[Req any, Resp any](ctx *httpx.Context, logicHandler LogicHandler[Req, Resp], b binding.Binding) (Resp, error) {
	var req Req
	var resp Resp
	if err := ctx.ShouldBindWith(&req, b); err != nil {
		if errors.Is(err, io.EOF) {
			logx.WithContext(ctx).Warn("ShouldBind异常", "io.EOF错误")
			return resp, err
		}
		logx.WithContext(ctx).Warn("ShouldBind异常", err)
		return resp, err
	}
	return logicHandler.Handle(ctx, req)
}
