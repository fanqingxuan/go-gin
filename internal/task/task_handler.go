package task

import (
	"github.com/hibiken/asynq"
)

type TaskHandler struct {
	typename string
	handler  asynq.HandlerFunc
}

func NewTaskHandler(typename string, handler asynq.HandlerFunc) *TaskHandler {
	return &TaskHandler{
		typename: typename,
		handler:  handler,
	}
}
