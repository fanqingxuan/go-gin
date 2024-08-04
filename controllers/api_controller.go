package controllers

import (
	"fmt"
	"go-gin/consts"
	"go-gin/internal/ginx/httpx"
	"go-gin/internal/httpc"

	"github.com/gin-gonic/gin"
)

type apiController struct {
}

var ApiController = &apiController{}

type User struct {
	Username string `json:"username"`
	Age      []int  `json:"age"`
}

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

func (c *apiController) Index(ctx *gin.Context) {

	var r Result
	_, err := httpc.POST(ctx, "http://localhost:8080/api/list").
		SetFormData(httpc.M{"username": "aaaa", "age": "55555"}).
		ParseResult(&r).
		Send()
	if err != nil {
		httpx.Error(ctx, consts.ErrThirdPartyAPIRequestFailed)
		return
	}
	fmt.Println(r)
	httpx.Ok(ctx, "ok")
}

func (c *apiController) IndexA(ctx *gin.Context) {

	var u User
	err := httpc.POST(ctx, "http://localhost:8080/api/list").
		SetFormData(httpc.M{"username": "aaaa", "age": "55555"}).
		SetHeader("hello", "测试").
		SendAndParseResult(&u)
	if err != nil {
		httpx.Error(ctx, consts.ErrThirdPartyAPIRequestFailed)
		return
	}
	httpx.Ok(ctx, u)
}

func (c *apiController) List(ctx *gin.Context) {
	httpx.Ok(ctx, gin.H{
		"username": ctx.PostForm("username"),
		"age":      ctx.PostFormArray("age"),
	})
}
