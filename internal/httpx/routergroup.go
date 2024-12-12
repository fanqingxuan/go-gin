package httpx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(*Context) (interface{}, error)
type MiddlewareFunc func(*Context)

// IRouter 定义路由接口
type IRouter interface {
	IRoutes
	Group(string, ...HandlerFunc) *RouterGroup
}

// IRoutes 定义路由方法接口
type IRoutes interface {
	Use(...HandlerFunc) IRoutes
	Handle(string, string, ...HandlerFunc) IRoutes
	Any(string, ...HandlerFunc) IRoutes
	GET(string, ...HandlerFunc) IRoutes
	POST(string, ...HandlerFunc) IRoutes
	DELETE(string, ...HandlerFunc) IRoutes
	PATCH(string, ...HandlerFunc) IRoutes
	PUT(string, ...HandlerFunc) IRoutes
	OPTIONS(string, ...HandlerFunc) IRoutes
	HEAD(string, ...HandlerFunc) IRoutes
}

// RouterGroup 包装 gin.RouterGroup
type RouterGroup struct {
	*gin.RouterGroup
}

// NewRouterGroup 创建包装后的路由组
func NewRouterGroup(group *gin.RouterGroup) *RouterGroup {
	return &RouterGroup{
		group,
	}
}

// Group 创建新的路由组
func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		group.RouterGroup.Group(relativePath, wrapHandlers(handlers)...),
	}
}

// Use 添加中间件
func (group *RouterGroup) Use(middleware ...HandlerFunc) IRoutes {
	group.RouterGroup.Use(wrapMiddlewares(middleware)...)
	return group
}

// Handle 注册请求处理函数
func (group *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) IRoutes {
	group.RouterGroup.Handle(httpMethod, relativePath, wrapHandlers(handlers)...)
	return group
}

// GET 处理 GET 请求
func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
	return group.Handle(http.MethodGet, relativePath, handlers...)
}

// POST 处理 POST 请求
func (group *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) IRoutes {
	return group.Handle(http.MethodPost, relativePath, handlers...)
}

// DELETE 处理 DELETE 请求
func (group *RouterGroup) DELETE(relativePath string, handlers ...HandlerFunc) IRoutes {
	return group.Handle(http.MethodDelete, relativePath, handlers...)
}

// PATCH 处理 PATCH 请求
func (group *RouterGroup) PATCH(relativePath string, handlers ...HandlerFunc) IRoutes {
	return group.Handle(http.MethodPatch, relativePath, handlers...)
}

// PUT 处理 PUT 请求
func (group *RouterGroup) PUT(relativePath string, handlers ...HandlerFunc) IRoutes {
	return group.Handle(http.MethodPut, relativePath, handlers...)
}

// OPTIONS 处理 OPTIONS 请求
func (group *RouterGroup) OPTIONS(relativePath string, handlers ...HandlerFunc) IRoutes {
	return group.Handle(http.MethodOptions, relativePath, handlers...)
}

// HEAD 处理 HEAD 请求
func (group *RouterGroup) HEAD(relativePath string, handlers ...HandlerFunc) IRoutes {
	return group.Handle(http.MethodHead, relativePath, handlers...)
}

// Any 处理任意 HTTP 方法
func (group *RouterGroup) Any(relativePath string, handlers ...HandlerFunc) IRoutes {
	group.GET(relativePath, handlers...)
	group.POST(relativePath, handlers...)
	group.PUT(relativePath, handlers...)
	group.PATCH(relativePath, handlers...)
	group.HEAD(relativePath, handlers...)
	group.OPTIONS(relativePath, handlers...)
	group.DELETE(relativePath, handlers...)
	return group
}

// wrapHandlers 包装处理函数
func wrapHandlers(handlers []HandlerFunc) []gin.HandlerFunc {
	return wrap(handlers, false)
}

// wrapMiddleware 包装中间件函数
func wrapMiddlewares(middlewares []HandlerFunc) []gin.HandlerFunc {
	return wrap(middlewares, true)
}

// wrap 包装处理函数
func wrap(handler []HandlerFunc, isMiddleware bool) []gin.HandlerFunc {
	wrapped := make([]gin.HandlerFunc, len(handler))
	for i, h := range handler {
		wrapped[i] = func(c *gin.Context) {
			ctx := NewContext(c)
			resp, err := h(ctx)
			if isMiddleware {
				if err != nil {
					Handle(ctx, resp, err)
					ctx.Abort()
					return
				}
				ctx.Next()
				return
			}
			Handle(ctx, resp, err)
		}
	}
	return wrapped
}
