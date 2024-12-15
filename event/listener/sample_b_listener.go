package listener

import (
	"context"
	"fmt"
	"go-gin/internal/eventbus"
)

type SampleBListener struct {
}

func (l *SampleBListener) Handle(ctx context.Context, e *eventbus.Event) error {
	fmt.Println("SampleBListener")
	return nil
}
