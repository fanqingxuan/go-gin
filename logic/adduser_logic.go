package logic

import (
	"context"
	"go-gin/internal/components/db"
	"go-gin/models"
	"go-gin/types"
)

type AddUserLogic struct {
}

func NewAddUserLogic() *AddUserLogic {
	return &AddUserLogic{}
}

func (l *AddUserLogic) Handle(ctx context.Context, req types.AddUserReq) (resp *types.AddUserReply, err error) {
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
