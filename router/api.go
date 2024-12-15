package router

import (
	"go-gin/controller"
	"go-gin/internal/httpx"
)

func RegisterApiRoutes(r *httpx.RouterGroup) {
	r.GET("/", controller.ApiController.Index)
	r.GET("/indexa", controller.ApiController.IndexA)
	r.GET("/loginapi", controller.ApiController.IndexB)
	r.GET("/mylogin", controller.ApiController.IndexC)
	r.Any("/list", controller.ApiController.List)
}
