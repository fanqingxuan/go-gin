package logic

import (
	"context"
	"go-gin/const/errcode"
	"go-gin/internal/token"
	"go-gin/model"
	"go-gin/typing"
)

type LoginLogic struct {
	model *model.UserModel
}

func NewLoginLogic() *LoginLogic {
	return &LoginLogic{
		model: model.NewUserModel(),
	}
}

func (l *LoginLogic) Handle(ctx context.Context, req typing.LoginReq) (resp *typing.LoginResp, err error) {
	if user, err := l.model.GetByUsername(ctx, req.Username); err != nil {
		if errcode.IsRecordNotFound(err) {
			return nil, errcode.ErrUserNotFound
		} else {
			return nil, err
		}
	} else {
		t := token.TokenId()
		if err := token.Set(ctx, t, "name", user.Name); err != nil {
			return nil, err
		}
		return &typing.LoginResp{
			Token: t,
			User:  *user,
		}, nil
	}

}
