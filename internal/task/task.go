package task

import "time"

type Task struct {
	taskName TaskName
	payload  any
}

type TaskName string

func (n *TaskName) Name() string {
	return string(*n)
}
func (t *Task) DispatchNow() error {
	return DispatchNow(t)
}

func (t *Task) Dispatch(d time.Duration) error {
	return Dispatch(t, d)
}

func NewTask(taskName TaskName, payload any) *Task {
	return &Task{
		taskName: taskName,
		payload:  payload,
	}
}
