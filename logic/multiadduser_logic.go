package logic

import (
	"context"
	"go-gin/model/dao"
	"go-gin/model/entity"
	"go-gin/typing"
)

type MultiAddUserLogic struct{}

func NewMultiAddUserLogic() *MultiAddUserLogic {
	return &MultiAddUserLogic{}
}

func (l *MultiAddUserLogic) Handle(ctx context.Context, req typing.MultiUserAddReq) (resp *typing.MultiUserAddResp, err error) {
	users := make([]*entity.User, len(req.Users))
	for i, user := range req.Users {
		users[i] = &entity.User{
			Name: user.Name,
		}
	}
	if err = dao.User.CreateBatch(ctx, users); err != nil {
		return nil, err
	}
	return &typing.MultiUserAddResp{
		Message: "批量添加成功",
	}, nil
}
