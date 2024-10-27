package token

import (
	"context"
	"fmt"
	"go-gin/internal/components/redisx"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	TOKEN_PREFIX = "token:"
	EXPIRE_TIME  = 7 * 24 * time.Hour
)

func TokenId() string {
	tokenId := uuid.New().String()
	return strings.ToLower(strings.ReplaceAll(tokenId, "-", ""))
}

func Set(ctx context.Context, key string, field string, value string) error {
	if err := redisx.GetInstance().HSet(ctx, transformKey(key), field, value).Err(); err != nil {
		return err
	}
	if err := redisx.GetInstance().Expire(ctx, transformKey(key), EXPIRE_TIME).Err(); err != nil {
		return err
	}
	return nil
}

func Get(ctx context.Context, key string, field string) (string, error) {
	cmd := redisx.GetInstance().HGet(ctx, transformKey(key), field)
	if err := cmd.Err(); err != nil {
		return "", err
	}
	return cmd.Val(), nil
}

func Has(ctx context.Context, key string) (bool, error) {
	cmd := redisx.GetInstance().Exists(ctx, transformKey(key))
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() != 0, nil
}

func HasField(ctx context.Context, key string, field string) (bool, error) {
	cmd := redisx.GetInstance().HExists(ctx, transformKey(key), field)
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val(), nil
}

func Delete(ctx context.Context, key string, field string) error {
	if err := redisx.GetInstance().HDel(ctx, transformKey(key), field).Err(); err != nil {
		return err
	}
	return nil
}

func Flush(ctx context.Context, key string) error {
	if err := redisx.GetInstance().Del(ctx, transformKey(key)).Err(); err != nil {
		return err
	}
	return nil
}

func GetAll(ctx context.Context, key string) (map[string]string, error) {
	cmd := redisx.GetInstance().HGetAll(ctx, transformKey(key))
	if err := cmd.Err(); err != nil {
		return nil, err
	}
	return cmd.Val(), nil
}

func transformKey(key string) string {
	return fmt.Sprintf("%s:%s", TOKEN_PREFIX, key)
}
