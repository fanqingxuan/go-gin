package controllers

import (
	"fmt"
	"go-gin/internal/ginx/httpx"
	"go-gin/internal/httpc"
	"go-gin/rest/user"

	"github.com/gin-gonic/gin"
)

type apiController struct {
}

var ApiController = &apiController{}

func (c *apiController) Index(ctx *gin.Context) {

	resp, err := httpc.POST(ctx, "http://localhost:8080/api/list").
		SetFormData(httpc.M{"username": "aaaa", "age": "55555"}).
		Send()
	if err != nil {
		httpx.Error(ctx, err)
		return
	}
	fmt.Println(resp)
	httpx.Ok(ctx, "ok")
}

func (c *apiController) IndexA(ctx *gin.Context) {

	resp, err := user.Svc.Hello(ctx, &user.HelloReq{UserId: "userId111"})
	if err != nil {
		httpx.Error(ctx, err)
		return
	}
	httpx.Ok(ctx, resp)
}

func (c *apiController) List(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "操作成功",
		"data": gin.H{
			"userId":   ctx.PostForm("userId"),
			"username": "张三",
			"age":      18,
		},
	})
}
