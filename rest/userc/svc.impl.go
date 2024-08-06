package userc

import (
	"context"
	"go-gin/internal/httpc"
)

type userSvc struct {
	client *httpc.Client
}

var _ IUserSvc = (*userSvc)(nil)

func NewUserSvc(client *httpc.Client) *userSvc {
	return &userSvc{
		client: client,
	}
}

func (us *userSvc) Hello(ctx context.Context, req *HelloReq) (resp *HelloResp, err error) {
	resp = &HelloResp{}
	d := APIResponse{Data: resp}
	err = us.client.
		NewRequest().
		SetContext(ctx).
		POST("/api/list").
		SetFormData(httpc.M{"username": "hello," + req.UserId, "age": "55555"}).
		SetResult(&d).
		Exec()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
