package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"go-gin/internal/task"
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
		func(ctx context.Context, data []byte) error {
			var p SampleBTaskPayload
			if err := json.Unmarshal(data, &p); err != nil {
				fmt.Println(err)
				return err
			}
			fmt.Println(p.UserId)
			// Image resizing code ...
			return nil
		})
}
