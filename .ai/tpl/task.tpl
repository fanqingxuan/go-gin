package task

import (
	"context"
	"fmt"
	"go-gin/internal/taskx"
)

const TypeXxxTask = "taskName"

func NewXxxTask(payload string) *queue.Task {
	return queue.NewTask(TypeXxxTask, payload)
}

func NewXxxTaskHandler() *queue.TaskHandler {
	return queue.NewTaskHandler(TypeXxxTask, func(ctx context.Context, data []byte) error {
		return nil
	})
}
