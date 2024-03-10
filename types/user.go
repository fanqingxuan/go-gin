package types

import "time"

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
	Name string `json:"name" form:"name"`
}

type User struct {
	Name  string `json:"name"`
	Ctime string `json:"ctime"`
}
type ListUserReply struct {
	User []User `json:"users"`
}
