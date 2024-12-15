package listener

import (
	"context"
	"fmt"
	"go-gin/internal/eventbus"
)

type SampleAListener struct {
}

func (l SampleAListener) Handle(ctx context.Context, e *eventbus.Event) error {
	fmt.Println("SampleAListener")
	fmt.Println(e.Payload().(string))
	return nil
}
