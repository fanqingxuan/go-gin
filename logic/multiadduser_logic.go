package logic

import (
	"context"
	"go-gin/model/dao"
	"go-gin/model/do"
	"go-gin/typing"
)

type MultiAddUserLogic struct{}

func NewMultiAddUserLogic() *MultiAddUserLogic {
	return &MultiAddUserLogic{}
}

func (l *MultiAddUserLogic) Handle(ctx context.Context, req typing.MultiUserAddReq) (resp *typing.MultiUserAddResp, err error) {
	users := make([]do.User, len(req.Users))
	for i, user := range req.Users {
		users[i] = do.User{Name: user.Name}
	}
	_, err = dao.User.Ctx(ctx).Data(users).Insert()
	if err != nil {
		return nil, err
	}
	return &typing.MultiUserAddResp{
		Message: "批量添加成功",
	}, nil
}
