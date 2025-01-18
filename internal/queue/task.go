package queue

import "time"

type Task struct {
	taskName string
	payload  any
}

func (t *Task) DispatchNow() error {
	return DispatchNow(t)
}

func (t *Task) DispatchIf(b bool) error {
	if !b {
		return nil
	}
	return DispatchNow(t)
}

func (t *Task) Dispatch(d time.Duration) error {
	return Dispatch(t, d)
}

func NewTask(taskName string, payload any) *Task {
	return &Task{
		taskName: taskName,
		payload:  payload,
	}
}
