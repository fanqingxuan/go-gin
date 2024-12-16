package task

import "go-gin/internal/taskx"

func Init() {
	taskx.AddHandler(NewSampleTaskHandler())
	taskx.AddHandler(NewSampleBTaskHandler())
}
