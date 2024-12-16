package task

import (
	"context"
	"encoding/json"
	"fmt"
	"go-gin/internal/taskx"
)

const TypeSampleBTask = "sampleB"

type SampleBTaskPayload struct {
	UserId []string
}

func NewSampleBTask(p string) *taskx.Task {
	return taskx.NewTask(TypeSampleBTask, SampleBTaskPayload{UserId: []string{p}})
}

func NewSampleBTaskHandler() *taskx.TaskHandler {
	return taskx.NewTaskHandler(TypeSampleBTask,
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
