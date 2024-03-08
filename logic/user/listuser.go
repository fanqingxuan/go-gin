package user

import (
	"go-gin/internal/errorx"
	"go-gin/types"
)

type ListUser struct {
}

func NewListUser() *ListUser {
	return &ListUser{}
}

func (this *ListUser) Handle(req types.ListUserReq) (*types.ListUserReply, error) {

	return nil, errorx.NewDefault("查询数据不存在")
}
