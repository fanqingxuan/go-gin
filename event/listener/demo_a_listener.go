package listener

import (
	"context"
	"fmt"
	"go-gin/internal/eventbus"
	"go-gin/model"
)

type DemoAListener struct {
}

func (l DemoAListener) Handle(ctx context.Context, e *eventbus.Event) error {
	user := e.Payload().(*model.User)
	fmt.Println(user.Name)
	return nil
}
