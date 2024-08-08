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
	Name   string    `form:"name" binding:"required"`
	Age    int       `form:"age" binding:"required" label:"年龄"`
	Status bool      `form:"status"`
	Ctime  time.Time `form:"ctime"`
}

type AddUserReply struct {
	Message string `json:"message"`
}
