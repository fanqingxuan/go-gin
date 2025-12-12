package logic

import (
	"context"
	"fmt"
	"go-gin/model/dao"
	"go-gin/model/do"
	"go-gin/typing"
)

type AddUserLogic struct{}

func NewAddUserLogic() *AddUserLogic {
	return &AddUserLogic{}
}

func (l *AddUserLogic) Handle(ctx context.Context, req typing.AddUserReq) (resp *typing.AddUserResp, err error) {
	id, err := dao.User.Ctx(ctx).Data(do.User{Name: req.Name}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	resp = &typing.AddUserResp{
		Message: fmt.Sprintf("message:%d", id),
	}
	return
}
