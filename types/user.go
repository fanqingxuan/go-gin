package types

import (
	"go-gin/models"
	"time"
)

type AddUserReq struct {
	Name   string    `json:"name" form:"name" binding:"required"`
	Age    int       `json:"age" form:"age" binding:"required"`
	Status bool      `json:"status" form:"status"`
	Ctime  time.Time `json:"ctime" form:"ctime"`
}

type AddUserReply struct {
	Message string `json:"message"`
}

type ListUserReq struct {
	Name string `json:"name" form:"name" binding:"required,email" label:"用户名"`
	Age  int    `json:"age" form:"age" binding:"min=0,max=100" label:"年龄"`
}

type ListUserReply struct {
	User []models.User `json:"users"`
}
