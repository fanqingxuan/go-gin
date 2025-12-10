package typing

import (
	"go-gin/model/entity"
	"time"
)

type ListReq struct {
}

type ListData struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	AgeTips      string `json:"age_tips"`
	Age          int    `json:"age"`
	Status       *int   `json:"status"`
	StatusText   string `json:"status_text"`
}

type ListResp struct {
	Data []ListData `json:"data"`
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
	Token string      `json:"token"`
	User  entity.User `json:"user"`
}

type UserData struct {
	Id   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Age  int    `json:"age" form:"age"`
}

type MultiUserAddReq struct {
	Users []UserData `json:"users" form:"users"`
}

type MultiUserAddResp struct {
	Message string `json:"message"`
}
