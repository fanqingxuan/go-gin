package logic

import (
	"context"
	"fmt"
	"go-gin/model/dao"
	"go-gin/model/entity"
	"go-gin/typing"
)

type AddUserLogic struct{}

func NewAddUserLogic() *AddUserLogic {
	return &AddUserLogic{}
}

func (l *AddUserLogic) Handle(ctx context.Context, req typing.AddUserReq) (resp *typing.AddUserResp, err error) {
	user := entity.User{
		Name: req.Name,
	}
	if err = dao.User.Create(ctx, &user); err != nil {
		return nil, err
	}
	resp = &typing.AddUserResp{
		Message: fmt.Sprintf("message:%d", user.Id),
	}
	return
}
