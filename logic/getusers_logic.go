package logic

import (
	"context"
	"go-gin/const/errcode"
	"go-gin/internal/components/redisx"
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
	var u []model.User
	if u, err = l.model.List(ctx); errcode.IsError(err) {
		return nil, err
	}

	redisx.Client().HSet(ctx, "name", "age", 43)

	return &typing.ListResp{
		Data: transformer.ConvertUserToListData(u),
	}, nil

}
