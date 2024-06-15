package services

import (
	"context"
	"fmt"
	"go-gin/internal/components/db"
	"go-gin/internal/components/redisx"
	"go-gin/models"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (svc *UserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var u []models.User
	if err := db.WithContext(ctx).Find(&u).Error; err != nil {
		return nil, err
	}

	redisx.GetInstance().HSet(ctx, "name", "age", 43)
	fmt.Println(redisx.GetInstance().Get(ctx, "name").String())
	return u, nil

}

func (svc *UserService) AddUser(ctx context.Context, user *models.User) error {
	if err := db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}
