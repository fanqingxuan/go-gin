package token

import (
	"context"
	"go-gin/internal/component/redisx"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	tokenPrefix = "token:"
	expireTime  = 7 * 24 * time.Hour
)

func TokenId() string {
	tokenId := uuid.New().String()
	return strings.ToLower(strings.ReplaceAll(tokenId, "-", ""))
}

func Set(ctx context.Context, key string, field string, value string) error {
	k := transformKey(key)
	if err := redisx.Client().HSet(ctx, k, field, value).Err(); err != nil {
		return err
	}
	return redisx.Client().Expire(ctx, k, expireTime).Err()
}

func Get(ctx context.Context, key string, field string) (string, error) {
	cmd := redisx.Client().HGet(ctx, transformKey(key), field)
	if err := cmd.Err(); err != nil {
		return "", err
	}
	return cmd.Val(), nil
}

func Has(ctx context.Context, key string) (bool, error) {
	cmd := redisx.Client().Exists(ctx, transformKey(key))
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() != 0, nil
}

func HasField(ctx context.Context, key string, field string) (bool, error) {
	cmd := redisx.Client().HExists(ctx, transformKey(key), field)
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val(), nil
}

func Delete(ctx context.Context, key string, field string) error {
	return redisx.Client().HDel(ctx, transformKey(key), field).Err()
}

func Flush(ctx context.Context, key string) error {
	return redisx.Client().Del(ctx, transformKey(key)).Err()
}

func GetAll(ctx context.Context, key string) (map[string]string, error) {
	cmd := redisx.Client().HGetAll(ctx, transformKey(key))
	if err := cmd.Err(); err != nil {
		return nil, err
	}
	return cmd.Val(), nil
}

func transformKey(key string) string {
	return tokenPrefix + key
}
