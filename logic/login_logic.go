package logic

import (
	"context"
	"go-gin/const/errcode"
	"go-gin/internal/token"
	"go-gin/model/dao"
	"go-gin/typing"
)

type LoginLogic struct{}

func NewLoginLogic() *LoginLogic {
	return &LoginLogic{}
}

func (l *LoginLogic) Handle(ctx context.Context, req typing.LoginReq) (resp *typing.LoginResp, err error) {
	user, err := dao.User.GetByName(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errcode.ErrUserNotFound
	}

	t := token.TokenId()
	if err := token.Set(ctx, t, "name", user.Name); err != nil {
		return nil, err
	}
	return &typing.LoginResp{
		Token: t,
		User:  *user,
	}, nil
}
