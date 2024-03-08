package user

import "go-gin/types"

type AddUser struct {
}

func NewAddUser() *AddUser {
	return &AddUser{}
}

func (this *AddUser) Handle(req types.AddUserReq) (*types.AddUserReply, error) {

	return &types.AddUserReply{
		Message: "hello," + req.Name,
	}, nil
}
