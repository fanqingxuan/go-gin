package listener

import (
	"context"
	"fmt"
	"go-gin/internal/eventbus"
)

type XxxBListener struct {
}

func (l *XxxBListener) Handle(ctx context.Context, e *eventbus.Event) error {
	return nil
}
