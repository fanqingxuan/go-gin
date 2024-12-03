package logic

import (
	"context"
	"go-gin/internal/components/redisx"
	"go-gin/models"
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

func (l *GetUsersLogic) Handle(ctx context.Context, req types.ListReq) (resp *types.ListReply, err error) {
	var u []models.User
	if u, err = l.model.List(ctx); err != nil {
		return nil, err
	}

	redisx.GetInstance().HSet(ctx, "name", "age", 43)
	return &types.ListReply{
		Users: u,
	}, nil

}
