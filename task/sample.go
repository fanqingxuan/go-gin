package task

import (
	"context"
	"fmt"
	"go-gin/internal/queue"
)

const TypeSampleTask = "sample"

func NewSampleTask(p string) *queue.Task {
	return queue.NewTask(TypeSampleTask, p)
}

func NewSampleTaskHandler() *queue.TaskHandler {
	return queue.NewTaskHandler(TypeSampleTask, func(ctx context.Context, data []byte) error {
		fmt.Println(string(data))
		// Image resizing code ...
		return nil
	})
}
