package types

import "time"

type AddUserReq struct {
	Name   string    `json:"name" form:"name"`
	Age    int       `json:"age" form:"age"`
	Status bool      `json:"status" form:"status"`
	Ctime  time.Time `json:"ctime",form:"ctime"`
}

type AddUserReply struct {
	Message string `json:"message"`
}
