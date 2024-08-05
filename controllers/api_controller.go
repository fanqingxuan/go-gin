package controllers

import (
	"fmt"
	"go-gin/consts"
	"go-gin/internal/ginx/httpx"
	"go-gin/internal/httpc"
	"go-gin/rest/userc"

	"github.com/gin-gonic/gin"
)

type apiController struct {
}

var ApiController = &apiController{}

type User struct {
	DDDD string `json:"username"`
	BBBB int
}

type Base struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type Result struct {
	Base
	Data User `json:"data"`
}

func (c *apiController) Index(ctx *gin.Context) {

	resp, err := httpc.POST(ctx, "http://localhost:8080/api/list").
		SetFormData(httpc.M{"username": "aaaa", "age": "55555"}).
		Send()
	if err != nil {
		httpx.Error(ctx, consts.ErrThirdPartyAPIRequestFailed)
		return
	}
	fmt.Println(resp)
	httpx.Ok(ctx, "ok")
}

func (c *apiController) IndexA(ctx *gin.Context) {

	var r Result
	err := httpc.POST(ctx, "http://localhost:8080/api/list").
		SetFormData(httpc.M{"username": "aaaa", "age": "55555"}).
		SetHeader("hello", "测试").
		SendAndParseResult(&r)
	if err != nil {
		httpx.Error(ctx, consts.ErrThirdPartyAPIRequestFailed)
		return
	}
	httpx.Ok(ctx, r)
}

func (c *apiController) IndexB(ctx *gin.Context) {

	resp, err := userc.UserSvc.Hello(ctx, &userc.HelloReq{UserId: "45"})
	if err != nil {
		httpx.Error(ctx, err)
		return
	}
	httpx.Ok(ctx, resp)
}

func (c *apiController) List(ctx *gin.Context) {
	httpx.Ok(ctx, gin.H{
		"username": ctx.PostForm("username"),
		"age":      4,
	})
}
