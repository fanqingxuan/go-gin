package mylogin

import (
	"context"
	"go-gin/internal/httpc"
)

const (
	Login_URL = "/logistics/apps/php/login.php?do=login#hello=json" // 登录接口
)

type LoginSvc struct {
	httpc.BaseSvc
}

func NewLoginSvc(url string) ILoginSvc {
	return &LoginSvc{
		BaseSvc: *httpc.NewBaseSvc(url),
	}
}

func (us *LoginSvc) Login(ctx context.Context, req *LoginReq) (resp *LoginResp, err error) {
	// md5加密

	params := httpc.M{"username": req.Username, "pwd": req.Pwd}
	result := APIResponse{Data: &resp}
	err = us.Client().
		NewRequest().
		SetContext(ctx).
		POST(Login_URL).
		SetFormData(params).
		SetResult(&result).
		Exec()
	return
}
