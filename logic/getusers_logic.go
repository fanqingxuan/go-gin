package logic

import (
	"context"
	"go-gin/internal/components/db"
	"go-gin/internal/components/redisx"
	"go-gin/models"
	"go-gin/types"
)

type GetUsersLogic struct {
}

func NewGetUsersLogic() *GetUsersLogic {
	return &GetUsersLogic{}
}

func (l *GetUsersLogic) Handle(ctx context.Context, req types.ListReq) (resp *types.ListReply, err error) {
	var u []models.User
	if err := db.WithContext(ctx).Find(&u).Error; err != nil {
		return nil, err
	}

	redisx.GetInstance().HSet(ctx, "name", "age", 43)
	return &types.ListReply{
		Users: u,
	}, nil

}
