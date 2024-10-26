package task

type Task struct {
	taskName TaskName
	payload  any
}

type TaskName string

func (n TaskName) Name() string {
	return string(n)
}

func NewTask(taskName TaskName, payload any) *Task {
	return &Task{
		taskName: taskName,
		payload:  payload,
	}
}
