package logic

import (
	"context"
	"go-gin/consts"
	"go-gin/internal/errorx"
	"go-gin/internal/token"
	"go-gin/models"
	"go-gin/types"
)

type LoginLogic struct {
	model *models.UserModel
}

func NewLoginLogic() *LoginLogic {
	return &LoginLogic{
		model: models.NewUserModel(),
	}
}

func (l *LoginLogic) Handle(ctx context.Context, req types.LoginReq) (resp *types.LoginReply, err error) {

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
		return &types.LoginReply{
			Token: t,
			User:  *user,
		}, nil
	}

}
