package redisx

import (
	"context"
	"go-gin/internal/errorx"
	"net"

	"github.com/redis/go-redis/v9"
)

type ErrHook struct{}

func (ErrHook) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		conn, err := next(ctx, network, addr)
		return conn, errorx.NewRedisError(err)
	}
}
func (ErrHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		return errorx.NewRedisError(next(ctx, cmd))
	}
}
func (ErrHook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		return errorx.NewRedisError(next(ctx, cmds))
	}
}
