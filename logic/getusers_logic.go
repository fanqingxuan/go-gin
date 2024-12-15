package logic

import (
	"context"
	"go-gin/internal/components/redisx"
	"go-gin/internal/errorx"
	"go-gin/models"
	"go-gin/transformers"
	"go-gin/types"
)

type GetUsersLogic struct {
	model *models.UserModel
}

func NewGetUsersLogic() *GetUsersLogic {
	return &GetUsersLogic{
		model: models.NewUserModel(),
	}
}

func (l *GetUsersLogic) Handle(ctx context.Context, req types.ListReq) (resp []types.ListResp, err error) {
	var u []models.User
	if u, err = l.model.List(ctx); errorx.IsError(err) {
		return nil, err
	}

	redisx.GetInstance().HSet(ctx, "name", "age", 43)

	return transformers.ConvertUserToListResp(u), nil

}
