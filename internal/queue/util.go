package queue

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/hibiken/asynq"
)

func DispatchNow(t *Task) error {
	return do(t)
}

func DispatchIf(b bool, t *Task) error {
	if !b {
		return nil
	}
	return DispatchNow(t)
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
	str := strings.Trim(string(p), "\"") // 结果: hello

	_, err = client.Enqueue(asynq.NewTask(t.taskName, []byte(str)), opts...)
	if err != nil {
		return err
	}
	return nil
}
