package taskx

import (
	"encoding/json"
	"time"

	"github.com/hibiken/asynq"
)

func DispatchNow(t *Task) error {
	return do(t)
}

func Dispatch(t *Task, d time.Duration) error {
	return do(t, asynq.ProcessIn(d))
}

func DispatchWithRetry(t *Task, d time.Duration, n int) error {
	return do(t, asynq.ProcessIn(d), asynq.MaxRetry(n))
}

func do(t *Task, opts ...asynq.Option) error {
	p, err := json.Marshal(t.payload)
	if err != nil {
		return err
	}
	_, err = client.Enqueue(asynq.NewTask(t.taskName, p), opts...)
	if err != nil {
		return err
	}
	return nil
}
