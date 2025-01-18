package httpx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Engine 包装 gin.Engine
type Engine struct {
	*gin.Engine
	RouterGroup // 继承 RouterGroup
}

// New 返回一个新的 Engine 实例
func New() *Engine {
	engine := &Engine{
		Engine: gin.New(),
	}
	engine.RouterGroup = *NewRouterGroup(&engine.Engine.RouterGroup)
	engine.HandleMethodNotAllowed = true
	return engine
}

// Default 返回一个带有默认中间件的 Engine 实例
func Default() *Engine {
	engine := New()
	engine.Use(recoverLog(), TraceId(), RequestLog(), dbCheck())
	return engine
}

// NoRoute 添加 404 处理器
func (engine *Engine) NoRoute(handlers ...HandlerFunc) {
	engine.Engine.NoRoute(wrapHandlers(handlers)...)
}

// NoMethod 添加 405 处理器
func (engine *Engine) NoMethod(handlers ...HandlerFunc) {
	engine.Engine.NoMethod(wrapHandlers(handlers)...)
}

// Group 创建一个新的路由组
func (engine *Engine) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return engine.RouterGroup.Group(relativePath, handlers...)
}

// Use 添加全局中间件
func (engine *Engine) Use(middleware ...gin.HandlerFunc) IRoutes {
	engine.RouterGroup.Use(middleware...)
	return engine
}

// Before 添加前置中间件
func (engine *Engine) Before(middleware ...HandlerFunc) IRoutes {
	engine.RouterGroup.Before(middleware...)
	return engine
}

// After 添加后置中间件
func (engine *Engine) After(middleware ...HandlerFunc) IRoutes {
	engine.RouterGroup.After(middleware...)
	return engine
}

// Routes 返回已注册的路由
func (engine *Engine) Routes() (routes gin.RoutesInfo) {
	return engine.Engine.Routes()
}

// Run 启动 HTTP 服务器
func (engine *Engine) Run(addr ...string) (err error) {
	return engine.Engine.Run(addr...)
}

// RunTLS 启动 HTTPS 服务器
func (engine *Engine) RunTLS(addr, certFile, keyFile string) (err error) {
	return engine.Engine.RunTLS(addr, certFile, keyFile)
}

// ServeHTTP 实现 http.Handler 接口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	engine.Engine.ServeHTTP(w, req)
}
