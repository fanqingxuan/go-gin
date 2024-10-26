package tasks

import "go-gin/internal/task"

func Init() {
	task.Handle(NewSampleTaskHandler())
	task.Handle(NewSampleBTaskHandler())
}
