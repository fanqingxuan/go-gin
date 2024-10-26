package task

type Task struct {
	typename string
	payload  any
}

func NewTask(typename string, payload any) *Task {
	return &Task{
		typename: typename,
		payload:  payload,
	}
}
