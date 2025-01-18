package cron

import (
	"context"
	"fmt"
)

func SampleFunc(ctx context.Context) error {
	fmt.Println("this is a sample function")
	return nil
}
