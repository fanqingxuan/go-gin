package logic

import (
	"context"
	"go-gin/const/errcode"
	"go-gin/model"
	"go-gin/transformer"
	"go-gin/typing"
)

type GetUsersLogic struct {
	model *model.UserModel
}

func NewGetUsersLogic() *GetUsersLogic {
	return &GetUsersLogic{
		model: model.NewUserModel(),
	}
}

func (l *GetUsersLogic) Handle(ctx context.Context, req typing.ListReq) (resp *typing.ListResp, err error) {
	u, err := l.model.List(ctx, "1=1")
	if errcode.IsError(err) {
		return nil, err
	}

	return &typing.ListResp{
		Data: transformer.ConvertUserToListData(u),
	}, nil
}
