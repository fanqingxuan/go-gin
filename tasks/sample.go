package tasks

import (
	"context"
	"fmt"
	"go-gin/internal/task"

	"github.com/hibiken/asynq"
)

const TypeSampleTask = "sample"

func NewSampleTask(p string) *task.Task {
	return task.NewTask(TypeSampleTask, p)
}

func NewSampleTaskHandler() *task.TaskHandler {
	return task.NewTaskHandler(TypeSampleTask, HandleSampleTask)
}

func HandleSampleTask(ctx context.Context, t *asynq.Task) error {
	fmt.Println(t.Type())
	fmt.Println(string(t.Payload()))
	// Image resizing code ...
	return nil
}
