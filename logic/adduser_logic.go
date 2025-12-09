package logic

import (
	"context"
	"fmt"
	"go-gin/const/enum"
	"go-gin/model"
	"go-gin/typing"
)

type AddUserLogic struct {
	model *model.UserModel
}

func NewAddUserLogic() *AddUserLogic {
	return &AddUserLogic{
		model: model.NewUserModel(),
	}
}

func (l *AddUserLogic) Handle(ctx context.Context, req typing.AddUserReq) (resp *typing.AddUserResp, err error) {
	user := model.User{
		Name:     req.Name,
		UserType: enum.UserTypeNormal,
	}
	if err = l.model.Create(ctx, &user); err != nil {
		return nil, err
	}
	resp = &typing.AddUserResp{
		Message: fmt.Sprintf("message:%d", user.Id),
	}
	return
}
