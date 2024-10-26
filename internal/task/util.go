package task

import (
	"encoding/json"
	"time"

	"github.com/hibiken/asynq"
)

func DispatchNow(t *Task) error {
	p, err := json.Marshal(t.payload)
	if err != nil {
		return err
	}
	_, err = client.Enqueue(asynq.NewTask(t.typename, p))
	if err != nil {
		return err
	}
	return nil
}

func Dispatch(t *Task, d time.Duration) error {
	p, err := json.Marshal(t.payload)
	if err != nil {
		return err
	}
	task := asynq.NewTask(t.typename, p)
	_, err = client.Enqueue(task, asynq.ProcessIn(d))
	if err != nil {
		return err
	}
	return nil
}

func DispatchWithRetry(t *Task, d time.Duration, n int) error {
	p, err := json.Marshal(t.payload)
	if err != nil {
		return err
	}
	task := asynq.NewTask(t.typename, p)
	_, err = client.Enqueue(task, asynq.ProcessIn(d), asynq.MaxRetry(n))
	if err != nil {
		return err
	}
	return nil
}
