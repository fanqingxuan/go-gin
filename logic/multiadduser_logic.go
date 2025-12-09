package logic

import (
	"context"
	"go-gin/model"
	"go-gin/typing"
)

type MultiAddUserLogic struct {
	model *model.UserModel
}

func NewMultiAddUserLogic() *MultiAddUserLogic {
	return &MultiAddUserLogic{
		model: model.NewUserModel(),
	}
}

func (l *MultiAddUserLogic) Handle(ctx context.Context, req typing.MultiUserAddReq) (resp *typing.MultiUserAddResp, err error) {
	users := make([]*model.User, len(req.Users))
	for i, user := range req.Users {
		users[i] = &model.User{
			Name: user.Name,
		}
	}
	if err = l.model.CreateBatch(ctx, users); err != nil {
		return nil, err
	}
	return &typing.MultiUserAddResp{
		Message: "批量添加成功",
	}, nil
}
