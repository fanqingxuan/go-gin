package user

import (
	"context"
	"go-gin/internal/httpc"
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
		POST("/api/list").
		SetFormData(params).
		SetResult(&result).
		Exec()
	return
}
