package logic

import (
	"context"
	"fmt"
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

	users := make([]model.User, len(req.Users))
	for i, user := range req.Users {
		users[i] = model.User{
			Name: user.Name,
		}
	}
	fmt.Println(users)
	return
}
