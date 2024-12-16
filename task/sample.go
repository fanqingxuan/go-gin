package task

import (
	"context"
	"fmt"
	"go-gin/internal/taskx"
)

const TypeSampleTask = "sample"

func NewSampleTask(p string) *taskx.Task {
	return taskx.NewTask(TypeSampleTask, p)
}

func NewSampleTaskHandler() *taskx.TaskHandler {
	return taskx.NewTaskHandler(TypeSampleTask, func(ctx context.Context, data []byte) error {
		fmt.Println(string(data))
		// Image resizing code ...
		return nil
	})
}
