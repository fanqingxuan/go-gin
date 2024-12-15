package types

import (
	"go-gin/model"
	"time"
)

type ListReq struct {
}

type ListResp struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	AgeTips string `json:"age_tips"`
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
	Username string `form:"username" binding:"required,email" label:"用户名"`
	Pwd      string `form:"pass" binding:"required,min=6" label:"密码"`
}

type LoginResp struct {
	Token string     `json:"token"`
	User  model.User `json:"user"`
}
