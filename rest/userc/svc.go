package userc

import "context"

type IUserSvc interface {
	Hello(context.Context, *HelloReq) (*HelloResp, error)
}

type HelloReq struct {
	UserId string
}

type HelloResp struct {
	Uname string `json:"username"`
}
