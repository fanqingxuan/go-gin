package listeners

import (
	"context"
	"fmt"
	"go-gin/internal/event"
)

type SampleBListener struct {
}

var _ event.Listener = (*SampleBListener)(nil)

func (l *SampleBListener) Handle(ctx context.Context, e *event.Event) error {
	fmt.Println("SampleBListener")
	return nil
}
