package task

import (
	"go-gin/internal/queue"
)

func Init() {
	queue.AddHandler(NewSampleTaskHandler())
	queue.AddHandler(NewSampleBTaskHandler())
}
