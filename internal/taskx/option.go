package taskx

import (
	"time"

	"github.com/hibiken/asynq"
)

type Option struct {
	opts []asynq.Option
}

func NewOption() *Option {
	return &Option{}
}

func (o *Option) MaxRetry(n int) *Option {
	o.add(asynq.MaxRetry(n))
	return o
}

func (o *Option) Queue(name string) *Option {
	o.add(asynq.Queue(name))
	return o
}

func (o *Option) TaskID(id string) *Option {
	o.add(asynq.TaskID(id))
	return o
}

func (o *Option) Timeout(d time.Duration) *Option {
	o.add(asynq.Timeout(d))
	return o
}

func (o *Option) Deadline(t time.Time) *Option {
	o.add(asynq.Deadline(t))
	return o
}

func (o *Option) Unique(ttl time.Duration) *Option {
	o.add(asynq.Unique(ttl))
	return o
}

func (o *Option) ProcessAt(t time.Time) *Option {
	o.add(asynq.ProcessAt(t))
	return o
}

func (o *Option) ProcessIn(d time.Duration) *Option {
	o.add(asynq.ProcessIn(d))
	return o
}

func (o *Option) Retention(d time.Duration) *Option {
	o.add(asynq.Retention(d))
	return o
}

func (o *Option) Group(name string) *Option {
	o.add(asynq.Group(name))
	return o
}

func (o *Option) Dispatch(t *Task) error {
	return do(t, o.opts...)
}

func (o *Option) add(opt asynq.Option) {
	o.opts = append(o.opts, opt)
}
