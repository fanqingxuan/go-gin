package task

import (
	"context"
	"encoding/json"
	"fmt"
	"go-gin/internal/queue"
)

const TypeSampleBTask = "sampleB"

type SampleBTaskPayload struct {
	UserId []string
}

func NewSampleBTask(p string) *queue.Task {
	return queue.NewTask(TypeSampleBTask, SampleBTaskPayload{UserId: []string{p}})
}

func NewSampleBTaskHandler() *queue.TaskHandler {
	return queue.NewTaskHandler(TypeSampleBTask,
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
