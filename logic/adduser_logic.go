package logic

import (
	"context"
	"fmt"
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
		Name: req.Name,
	}
	if err = l.model.Add(ctx, &user); err != nil {
		return
	}
	resp = &typing.AddUserResp{
		Message: fmt.Sprintf("message:%d", user.Id),
	}
	return
}
