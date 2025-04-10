package redisx

import (
	"context"
	"go-gin/internal/component/logx"
	"go-gin/internal/util"
	"net"

	"github.com/redis/go-redis/v9"
)

type LogHook struct{}

func (LogHook) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		return next(ctx, network, addr)

	}
}
func (LogHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		err := next(ctx, cmd)
		if err != nil {
			logx.WithContext(ctx).Warnf("redis", "%s execute command:%+v, error=%s", util.FileWithLineNum(), cmd.Args(), err)
			return err
		}
		return nil
	}
}
func (LogHook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		return next(ctx, cmds)
	}
}
