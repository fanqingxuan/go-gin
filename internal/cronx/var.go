package cronx

import "context"

type Job interface {
	Handle(ctx context.Context) error
}

type JobFunc func(context.Context) error

var _ Job = JobFunc(nil)

func (f JobFunc) Handle(ctx context.Context) error {
	return f(ctx)
}
