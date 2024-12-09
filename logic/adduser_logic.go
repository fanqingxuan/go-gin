package logic

import (
	"context"
	"fmt"
	"go-gin/models"
	"go-gin/types"
)

type AddUserLogic struct {
	model *models.UserModel
}

func NewAddUserLogic() *AddUserLogic {
	return &AddUserLogic{
		model: models.NewUserModel(),
	}
}

func (l *AddUserLogic) Handle(ctx context.Context, req types.AddUserReq) (resp *types.AddUserResp, err error) {
	user := models.User{
		Name: req.Name,
	}
	if err = l.model.Add(ctx, &user); err != nil {
		return
	}
	resp = &types.AddUserResp{
		Message: fmt.Sprintf("message:%d", user.Id),
	}
	return
}
