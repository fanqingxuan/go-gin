package controller

import (
	"fmt"
	"go-gin/internal/httpx"
	"go-gin/logic"
	"go-gin/types"

	"github.com/golang-module/carbon/v2"
)

type userController struct {
}

var UserController = &userController{}

type User struct {
	Name       string        `json:"name"`
	CreateTime carbon.Carbon `json:"create_time"`
}

func (c *userController) Index(ctx *httpx.Context) (interface{}, error) {
	// event.Fire(ctx, events.NewSampleEvent("hello 测试"))
	// events.NewSampleEvent("333").Fire(ctx)
	// u := User{
	// 	Name:       "hello",
	// 	CreateTime: carbon.Parse("now").AddCentury(),
	// }
	// return u, nil
	fmt.Println("index")
	return ShouldBindHandle(ctx, logic.NewIndexLogic())
}

func (c *userController) List(ctx *httpx.Context) (interface{}, error) {
	var req types.ListReq
	l := logic.NewGetUsersLogic()
	return l.Handle(ctx, req)
}

func (c *userController) AddUser(ctx *httpx.Context) (interface{}, error) {
	return ShouldBindHandle(ctx, logic.NewAddUserLogic())
}
