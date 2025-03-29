package logic

import (
	"context"
	"fmt"
	"go-gin/const/errcode"
	"go-gin/internal/components/redisx"
	"go-gin/model"
	"go-gin/transformer"
	"go-gin/typing"
	"go-gin/util/jsonx"
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

	jsonStr, err := jsonx.Encode(u)
	fmt.Println(jsonStr, err)

	var my []model.User
	err = jsonx.Decode(jsonStr, &my)
	fmt.Println(my, err)

	redisx.Client().HSet(ctx, "name", "age", 43)

	return &typing.ListResp{
		Data: transformer.ConvertUserToListData(u),
	}, nil

}
