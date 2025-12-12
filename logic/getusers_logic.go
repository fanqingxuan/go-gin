package logic

import (
	"context"
	"go-gin/const/errcode"
	"go-gin/model/dao"
	"go-gin/model/entity"
	"go-gin/transformer"
	"go-gin/typing"
)

type GetUsersLogic struct{}

func NewGetUsersLogic() *GetUsersLogic {
	return &GetUsersLogic{}
}

func (l *GetUsersLogic) Handle(ctx context.Context, req typing.ListReq) (resp *typing.ListResp, err error) {
	var users []*entity.User
	err = dao.User.Ctx(ctx).All(&users)
	if errcode.IsError(err) {
		return nil, err
	}

	return &typing.ListResp{
		Data: transformer.ConvertUserToListData(users),
	}, nil
}
