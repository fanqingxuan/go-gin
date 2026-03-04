# Controller 控制器层

`controller/` 目录是 HTTP 请求的入口层，职责是接收请求、委托 Logic 处理、返回响应。Controller 是薄层，禁止编写业务逻辑。

## 基本结构

每个 Controller 包含：unexported 结构体、exported 单例、处理方法。

```go
// controller/user_controller.go
package controller

import (
    "go-gin/internal/httpx"
    "go-gin/logic"
)

type userController struct{}

var UserController = &userController{}

func (c *userController) AddUser(ctx *httpx.Context) (any, error) {
    return httpx.Handle(ctx, logic.NewAddUserLogic())
}
```

## Handler 签名

所有 Controller 方法统一签名，框架自动处理 JSON 响应：

```go
func (c *xxxController) MethodName(ctx *httpx.Context) (any, error)
```

- 返回 `(data, nil)` → 成功响应 `{"code": 200, "data": ...}`
- 返回 `(nil, err)` → 错误响应（根据错误类型自动判断 HTTP 状态码和 code）
- 返回 `(nil, nil)` → 成功响应无 data

## 请求绑定方法

通过 `httpx.Handle` 系列函数自动绑定请求参数并调用 Logic：

| 方法 | 说明 | 适用场景 |
|------|------|---------|
| `httpx.Handle(ctx, logic)` | 自动根据 Content-Type 绑定 | 通用（推荐） |
| `httpx.HandleJSON(ctx, logic)` | 强制 JSON 绑定 | 明确要求 JSON body |
| `httpx.HandleQuery(ctx, logic)` | Query 参数绑定 | GET 请求参数 |
| `httpx.HandleUri(ctx, logic)` | URI 参数绑定 | 路径参数 `/user/:id` |
| `httpx.HandleHeader(ctx, logic)` | Header 绑定 | 从请求头提取参数 |
| `httpx.HandleWith(ctx, logic, binding)` | 指定 Binding | 自定义绑定方式 |

## 使用模式

### 模式 1：标准 Logic 委托（推荐）

```go
func (c *userController) List(ctx *httpx.Context) (any, error) {
    return httpx.Handle(ctx, logic.NewGetUsersLogic())
}

func (c *userController) AddUser(ctx *httpx.Context) (any, error) {
    return httpx.Handle(ctx, logic.NewAddUserLogic())
}
```

### 模式 2：无需参数绑定的简单操作

```go
func (c *loginController) LoginOut(ctx *httpx.Context) (any, error) {
    token.Flush(ctx, ctx.GetHeader("token"))
    return nil, nil
}
```

### 模式 3：直接返回数据

```go
func (c *apiController) List(ctx *httpx.Context) (any, error) {
    return gin.H{
        "userId":   ctx.PostForm("userId"),
        "username": "张三",
        "age":      18,
    }, nil
}
```

### 模式 4：调用 REST 服务

```go
func (c *apiController) IndexA(ctx *httpx.Context) (any, error) {
    return user.Svc.Hello(ctx, &user.HelloReq{UserId: "123"})
}
```

### 模式 5：使用 HTTP 客户端

```go
func (c *apiController) Index(ctx *httpx.Context) (any, error) {
    return httpc.POST(ctx, "http://localhost:8080/api/list").
        SetFormData(httpc.M{"username": "aaaa", "age": "55555"}).
        Send()
}
```

## 路由注册

Controller 方法在 `router/` 中注册：

```go
// router/user.go
func RegisterUserRoutes(r *httpx.RouterGroup) {
    r.GET("/index", controller.UserController.Index)
    r.GET("/list", controller.UserController.List)
    r.POST("/add", controller.UserController.AddUser)
}
```

支持路由组级别中间件：

```go
r := route.Group("/user")
r.Before(middleware.TokenCheck())
r.GET("/profile", controller.UserController.Profile)
```

## 请求类型定义（typing/）

请求结构体使用 `form:` tag 绑定参数，`binding:` tag 校验，`label:` tag 提供中文错误提示：

```go
// typing/user.go
type AddUserReq struct {
    Name   string    `form:"name" binding:"required" label:"姓名"`
    Age    int       `form:"age" binding:"required" label:"年龄"`
    Status bool      `form:"status"`
    Ctime  time.Time `form:"ctime"`
}

type LoginReq struct {
    Username string `form:"username" binding:"required,email" label:"用户名"`
    Pwd      string `form:"pass" binding:"required,min=6" label:"密码"`
}
```

## 命名规范

| 项目 | 规范 | 示例 |
|------|------|------|
| 文件名 | `{module}_controller.go` | `user_controller.go`、`login_controller.go` |
| 结构体（unexported） | `{module}Controller` | `userController`、`loginController` |
| 单例（exported） | `{Module}Controller` | `UserController`、`LoginController` |
| 方法名 | PascalCase 动词 | `List`、`AddUser`、`Login` |

## 禁止事项

1. **禁止在 Controller 中写业务逻辑** —— 委托给 Logic 层
2. **禁止在 Controller 中直接操作 db** —— 使用 `dao.Xxx.Ctx(ctx)` 链式调用
3. **禁止手动绑定参数 + 调用 Logic** —— 使用 `httpx.Handle` 统一处理
4. **禁止手动构造 JSON 响应** —— 框架自动处理

```go
// ❌ 避免手动绑定
func (c *userController) AddUser(ctx *httpx.Context) (any, error) {
    var req typing.AddUserReq
    if err := ctx.ShouldBind(&req); err != nil { ... }
    return logic.NewAddUserLogic().Handle(ctx, req)
}

// ✅ 推荐
func (c *userController) AddUser(ctx *httpx.Context) (any, error) {
    return httpx.Handle(ctx, logic.NewAddUserLogic())
}
```
