package listeners

import (
	"context"
	"fmt"
	"go-gin/internal/event"
)

type SampleAListener struct {
}

var _ event.Listener = (*SampleAListener)(nil)

func (l SampleAListener) Handle(ctx context.Context, e event.Event) error {
	fmt.Println("SampleAListener")
	fmt.Println(e.Name())
	fmt.Println(e.Payload().(string))
	return nil
}
