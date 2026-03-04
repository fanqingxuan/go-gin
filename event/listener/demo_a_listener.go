package listener

import (
	"context"
	"fmt"
	"go-gin/internal/eventbus"
	"go-gin/model/entity"
)

type DemoAListener struct {
}

func (l *DemoAListener) Handle(ctx context.Context, e *eventbus.Event) error {
	user := eventbus.PayloadAs[*entity.User](e)
	fmt.Println(user.Name)
	return nil
}
