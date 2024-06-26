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

func GetTokenId() string {
	tokenId := uuid.New().String()
	return strings.ToLower(strings.ReplaceAll(tokenId, "-", ""))
}

func Set(ctx context.Context, key string, field string, value string) error {
	fmt.Println(redisx.GetInstance().HSet(ctx, getKey(key), field, value).Err())
	if err := redisx.GetInstance().HSet(ctx, getKey(key), field, value).Err(); err != nil {
		return err
	}
	if err := redisx.GetInstance().Expire(ctx, getKey(key), EXPIRE_TIME).Err(); err != nil {
		return err
	}
	return nil
}

func Get(ctx context.Context, key string, field string) (string, error) {
	cmd := redisx.GetInstance().HGet(ctx, getKey(key), field)
	if err := cmd.Err(); err != nil {
		return "", err
	}
	return cmd.Val(), nil
}

func Del(ctx context.Context, key string, field string) error {
	if err := redisx.GetInstance().HDel(ctx, getKey(key), field).Err(); err != nil {
		return err
	}
	return nil
}

func FlushAll(ctx context.Context, key string) error {
	if err := redisx.GetInstance().Del(ctx, getKey(key)).Err(); err != nil {
		return err
	}
	return nil
}

func GetAll(ctx context.Context, key string) (map[string]string, error) {
	cmd := redisx.GetInstance().HGetAll(ctx, getKey(key))
	if err := cmd.Err(); err != nil {
		return nil, err
	}
	return cmd.Val(), nil
}

func getKey(key string) string {
	return fmt.Sprintf("%s:%s", TOKEN_PREFIX, key)
}
