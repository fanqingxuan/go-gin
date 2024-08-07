package user

import "context"

type IUserSvc interface {
	Hello(context.Context, *HelloReq) (*HelloResp, error)
}

type HelloReq struct {
	UserId string
}

type HelloResp struct {
	Uid   string `json:"userId"`
	Uname string `json:"username"`
}
