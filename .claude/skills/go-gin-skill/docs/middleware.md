# Middleware 中间件

中间件定义在 `middleware/` 目录，在 `middleware/init.go` 中注册。

## 中间件签名

使用 `httpx.HandlerFunc`：`func(*httpx.Context) (any, error)`

返回 error 时自动中止请求并返回错误响应。

## 全局中间件 (middleware/init.go)

```go
package middleware

import "go-gin/internal/httpx"

func Init(r *httpx.Engine) {
    r.Before(BeforeSampleA(), BeforeSampleB())  // 前置中间件
    r.After(AfterSampleB())                      // 后置中间件
    // r.Before(TokenCheck())                    // Token 校验
}
```

## Before 中间件（请求前执行）

```go
// middleware/before_sample_a.go
package middleware

import "go-gin/internal/httpx"

func BeforeSampleA() httpx.HandlerFunc {
    return func(ctx *httpx.Context) (any, error) {
        // 返回 nil, nil 表示通过，继续执行
        return nil, nil
    }
}
```

## After 中间件（请求后执行）

```go
// middleware/after_sample_a.go
func AfterSampleA() httpx.HandlerFunc {
    return func(ctx *httpx.Context) (any, error) {
        // 请求处理完成后执行
        return nil, nil
    }
}
```

## Token 校验中间件示例

```go
// middleware/token_check.go
func TokenCheck() httpx.HandlerFunc {
    return func(ctx *httpx.Context) (any, error) {
        var req TokenHeader
        if err := ctx.ShouldBindHeader(&req); err != nil {
            return nil, errcode.ErrUserMustLogin
        }
        if has, err := token.Has(ctx, req.Token); err != nil {
            return nil, errcode.NewDefault("获取token错误")
        } else if !has {
            return nil, errcode.ErrUserNeedLoginAgain
        }
        return nil, nil
    }
}
```

## 路由组级别中间件

```go
// 在路由注册时使用
r := route.Group("/user")
r.Before(TokenCheck())  // 仅对 /user 下路由生效
r.GET("/profile", controller.UserController.Profile)
```
