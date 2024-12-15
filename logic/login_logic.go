package logic

import (
	"context"
	"go-gin/consts"
	"go-gin/internal/errorx"
	"go-gin/internal/token"
	"go-gin/model"
	"go-gin/types"
)

type LoginLogic struct {
	model *model.UserModel
}

func NewLoginLogic() *LoginLogic {
	return &LoginLogic{
		model: model.NewUserModel(),
	}
}

func (l *LoginLogic) Handle(ctx context.Context, req types.LoginReq) (resp *types.LoginResp, err error) {
	if user, err := l.model.GetByUsername(ctx, req.Username); err != nil {
		if errorx.IsRecordNotFound(err) {
			return nil, consts.ErrUserNotFound
		} else {
			return nil, err
		}
	} else {
		t := token.TokenId()
		if err := token.Set(ctx, t, "name", user.Name); err != nil {
			return nil, err
		}
		return &types.LoginResp{
			Token: t,
			User:  *user,
		}, nil
	}

}
