package errorx

import (
	"errors"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func IsRecordNotFound(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, redis.Nil)
}
