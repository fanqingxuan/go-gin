package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"go-gin/internal/task"

	"github.com/hibiken/asynq"
)

const TypeSampleBTask = "sampleB"

type SampleBTaskPayload struct {
	UserId []string
}

func NewSampleBTask(p string) *task.Task {
	return task.NewTask(TypeSampleBTask, SampleBTaskPayload{UserId: []string{p}})
}

func NewSampleBTaskHandler() *task.TaskHandler {
	return task.NewTaskHandler(TypeSampleBTask,
		func(ctx context.Context, t *asynq.Task) error {
			fmt.Println(t.Type())
			var p SampleBTaskPayload
			if err := json.Unmarshal(t.Payload(), &p); err != nil {
				fmt.Println(err)
				return err
			}
			fmt.Println(p.UserId)
			// Image resizing code ...
			return nil
		})
}
