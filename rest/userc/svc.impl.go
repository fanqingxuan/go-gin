package userc

import (
	"context"
	"fmt"
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
	err = us.client.
		NewRequest().
		SetContext(ctx).
		POST("/api/list").
		SetFormData(httpc.M{"username": "hello," + req.UserId, "age": "55555"}).
		SendAndParseResult(WrapUserStruct(&resp))
	fmt.Println(resp.Uname)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
