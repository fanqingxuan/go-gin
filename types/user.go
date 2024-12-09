package types

import (
	"go-gin/models"
	"time"
)

type ListReq struct {
}

type ListReply struct {
	Users []models.User `json:"users"`
}

type AddUserReq struct {
	Name   string    `form:"name" binding:"required" label:"姓名"`
	Age    int       `form:"age" binding:"required" label:"年龄"`
	Status bool      `form:"status"`
	Ctime  time.Time `form:"ctime"`
}

type AddUserResp struct {
	Message string `json:"message"`
}

type LoginReq struct {
	Username string `form:"username" binding:"required" label:"用户名"`
}

type LoginResp struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}
