package user

import (
	"context"
	"go-gin/internal/httpc"
)

const (
	HELLO_URL = "/api/list" // hello的接口路径
)

type UserSvc struct {
	httpc.BaseSvc
}

func NewUserSvc(url string) IUserSvc {
	return &UserSvc{
		BaseSvc: *httpc.NewBaseSvc(url),
	}
}

func (us *UserSvc) Hello(ctx context.Context, req *HelloReq) (resp *HelloResp, err error) {
	params := httpc.M{"userId": req.UserId}
	result := APIResponse{Data: &resp}
	err = us.Client().
		NewRequest().
		SetContext(ctx).
		POST(HELLO_URL).
		SetFormData(params).
		SetResult(&result).
		Exec()
	return
}
