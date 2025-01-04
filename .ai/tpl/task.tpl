package task

import (
	"context"
	"fmt"
	"go-gin/internal/taskx"
)

const TypeXxxTask = "taskName"

func NewXxxTask(payload string) *taskx.Task {
	return taskx.NewTask(TypeXxxTask, payload)
}

func NewXxxTaskHandler() *taskx.TaskHandler {
	return taskx.NewTaskHandler(TypeXxxTask, func(ctx context.Context, data []byte) error {
		return nil
	})
}
