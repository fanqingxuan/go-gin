package cron

import "context"

type Job interface {
	Name() string
	Handle(ctx context.Context) error
}
