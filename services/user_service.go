package services

import (
	"context"
	"go-gin/internal/components/db"
	"go-gin/internal/components/redisx"
	"go-gin/models"
	"go-gin/types"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (svc *UserService) GetAllUsers(ctx context.Context, req types.ListReq) (resp *types.ListReply, err error) {
	var u []models.User
	if err := db.WithContext(ctx).Find(&u).Error; err != nil {
		return nil, err
	}

	redisx.GetInstance().HSet(ctx, "name", "age", 43)
	return &types.ListReply{
		Users: u,
	}, nil

}

func (svc *UserService) AddUser(ctx context.Context, req types.AddUserReq) (resp *types.AddUserReply, err error) {
	user := models.User{
		Name: req.Name,
	}
	if err = db.WithContext(ctx).Select("Name").Create(&user).Error; err != nil {
		return
	}
	resp = &types.AddUserReply{
		Message: "success",
	}
	return
}
