package httpx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(*Context) (interface{}, error)

// IRouter 定义路由接口
type IRouter interface {
	IRoutes
	Group(string, ...HandlerFunc) *RouterGroup
}

// IRoutes 定义路由方法接口
type IRoutes interface {
	Use(...gin.HandlerFunc) IRoutes
	Before(...HandlerFunc) IRoutes
	After(...HandlerFunc) IRoutes
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
func (group *RouterGroup) Use(middleware ...gin.HandlerFunc) IRoutes {
	group.RouterGroup.Use(middleware...)
	return group
}

// Before 添加前置中间件
func (group *RouterGroup) Before(middleware ...HandlerFunc) IRoutes {
	wrapped := make([]gin.HandlerFunc, len(middleware))
	for i, h := range middleware {
		wrapped[i] = func(m HandlerFunc) gin.HandlerFunc {
			return func(c *gin.Context) {
				ctx := NewContext(c)
				_, err := m(ctx)
				if err != nil {
					Handle(ctx, nil, err)
					c.Abort()
					return
				}
				c.Next()
			}
		}(h)
	}
	group.RouterGroup.Use(wrapped...)
	return group
}

// After 添加后置中间件
func (group *RouterGroup) After(middleware ...HandlerFunc) IRoutes {
	wrapped := make([]gin.HandlerFunc, len(middleware))
	for i, h := range middleware {
		wrapped[i] = func(m HandlerFunc) gin.HandlerFunc {
			return func(c *gin.Context) {
				c.Next()
				m(NewContext(c))
			}
		}(h)
	}
	group.RouterGroup.Use(wrapped...)
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

// wrap 包装处理函数
func wrapHandlers(handler []HandlerFunc) []gin.HandlerFunc {
	wrapped := make([]gin.HandlerFunc, len(handler))
	for i, h := range handler {
		wrapped[i] = func(c *gin.Context) {
			ctx := NewContext(c)
			resp, err := h(ctx)
			Handle(ctx, resp, err)
			c.Abort()
		}
	}
	return wrapped
}
