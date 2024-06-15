package types

import (
	"time"
)

type AddUserReq struct {
	Name   string    `form:"name" binding:"required"`
	Age    int       `form:"age" binding:"required" label:"年龄"`
	Status bool      `form:"status"`
	Ctime  time.Time `form:"ctime"`
}

type AddUserReply struct {
	Message string `json:"message"`
}
