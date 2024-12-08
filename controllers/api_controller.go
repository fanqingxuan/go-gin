package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"go-gin/internal/httpc"
	"go-gin/internal/httpx"
	"go-gin/rest/login"
	"go-gin/rest/mylogin"
	"go-gin/rest/user"

	"github.com/gin-gonic/gin"
)

type apiController struct {
}

var ApiController = &apiController{}

func (c *apiController) Index(ctx *httpx.Context) (interface{}, error) {

	return httpc.POST(ctx, "http://localhost:8080/api/list").
		SetFormData(httpc.M{"username": "aaaa", "age": "55555"}).
		Send()

}

func (c *apiController) IndexA(ctx *httpx.Context) (interface{}, error) {

	return user.Svc.Hello(ctx, &user.HelloReq{UserId: "userId111"})

}

func (c *apiController) IndexB(ctx *httpx.Context) (interface{}, error) {

	hash := md5.Sum([]byte("BRUCEMUWU2023"))
	pwd := hex.EncodeToString(hash[:])
	return login.Svc.Login(ctx, &login.LoginReq{Username: "1", Pwd: pwd})
}

func (c *apiController) IndexC(ctx *httpx.Context) (interface{}, error) {
	hash := md5.Sum([]byte("BRUCEMUWU2"))
	pwd := hex.EncodeToString(hash[:])
	return mylogin.Svc.Login(ctx, &mylogin.LoginReq{Username: "1", Pwd: pwd})
}

func (c *apiController) List(ctx *httpx.Context) (interface{}, error) {
	return gin.H{
		"userId":   ctx.PostForm("userId"),
		"username": "张三",
		"age":      18,
	}, nil
}
