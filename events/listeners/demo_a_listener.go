package listeners

import (
	"context"
	"fmt"
	"go-gin/internal/event"
	"go-gin/models"
)

type DemoAListener struct {
}

func (l DemoAListener) Handle(ctx context.Context, e *event.Event) error {
	user := e.Payload().(*models.User)
	fmt.Println(user.Name)
	return nil
}
