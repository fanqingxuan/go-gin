package tasks

import "go-gin/internal/task"

func Init() {
	task.AddHandler(NewSampleTaskHandler())
	task.AddHandler(NewSampleBTaskHandler())
}
