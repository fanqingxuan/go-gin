package controllers

import (
	"go-gin/consts"
	"go-gin/internal/ginx/httpx"
	"go-gin/internal/request"

	"github.com/gin-gonic/gin"
)

type apiController struct {
}

var ApiController = &apiController{}

type User struct {
	Code int
	Data struct {
		Username string `json:"username"`
	}
}

func (c *apiController) Index(ctx *gin.Context) {

	var u User
	_, err := request.POST(ctx, "http://localhosft:8080/api/list").
		SetFormData(request.M{"username": "aaaa", "age": "55555"}).
		ParseResult(&u).
		Send()
	if err != nil {
		httpx.Error(ctx, consts.ErrThirdPartyAPIRequestFailed)
		return
	}
	httpx.Ok(ctx, "ok")
}

func (c *apiController) List(ctx *gin.Context) {
	httpx.Ok(ctx, gin.H{
		"username": ctx.PostForm("username"),
		"age":      ctx.PostFormArray("age"),
	})
}
