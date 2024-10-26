package tasks

import (
	"context"
	"fmt"
	"go-gin/internal/task"
)

const TypeSampleTask = "sample"

func NewSampleTask(p string) *task.Task {
	return task.NewTask(TypeSampleTask, p)
}

func NewSampleTaskHandler() *task.TaskHandler {
	return task.NewTaskHandler(TypeSampleTask, func(ctx context.Context, data []byte) error {
		fmt.Println(string(data))
		// Image resizing code ...
		return nil
	})
}
