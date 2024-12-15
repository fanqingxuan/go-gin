package user

import (
	"context"
	"go-gin/internal/httpc"
	"go-gin/util/jsonx"
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
	result := APIResponse{Data: &resp}
	ids, _ := jsonx.MarshalToString([]int{1, 2, 3})

	userInfo, _ := jsonx.MarshalToString(httpc.M{
		"id":   "3",
		"name": "张三",
	})
	params := httpc.M{"userId": req.UserId, "ids": ids, "info": userInfo}

	err = us.Client().
		NewRequest().
		SetContext(ctx).
		POST(HELLO_URL).
		SetFormData(params).
		AddFormData("userA", []string{"11", "22"}).
		SetResult(&result).
		Exec()
	return
}
