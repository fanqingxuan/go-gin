package listeners

import (
	"context"
	"fmt"
	"go-gin/internal/event"
	"go-gin/model"
)

type DemoAListener struct {
}

func (l DemoAListener) Handle(ctx context.Context, e *event.Event) error {
	user := e.Payload().(*model.User)
	fmt.Println(user.Name)
	return nil
}
