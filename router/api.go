package router

import (
	"go-gin/controllers"
	"go-gin/internal/httpx"
)

func RegisterApiRoutes(r *httpx.RouterGroup) {
	r.GET("/", controllers.ApiController.Index)
	r.GET("/indexa", controllers.ApiController.IndexA)
	r.GET("/loginapi", controllers.ApiController.IndexB)
	r.GET("/mylogin", controllers.ApiController.IndexC)
	r.Any("/list", controllers.ApiController.List)
}
