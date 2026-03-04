# REST 第三方服务调用

`rest/` 目录封装第三方 HTTP API 调用，每个服务一个子目录。

## 目录结构

```
rest/
└── user/
    ├── init.go         # 初始化，暴露 Svc 单例
    ├── svc.go          # 接口定义 + 请求/响应类型
    ├── svc.impl.go     # 接口实现（HTTP 调用）
    └── response.go     # 响应结构解析
```

## 1. 定义接口和类型 (svc.go)

```go
package user

import "context"

type IUserSvc interface {
    Hello(context.Context, *HelloReq) (*HelloResp, error)
}

type HelloReq struct {
    UserId string
}

type HelloResp struct {
    Uid   string `json:"userId"`
    Uname string `json:"username"`
}
```

## 2. 实现接口 (svc.impl.go)

```go
package user

import (
    "context"
    "go-gin/internal/httpc"
)

const helloURL = "/api/list"

type UserSvc struct {
    httpc.BaseSvc
}

func NewUserSvc(url string) IUserSvc {
    return &UserSvc{BaseSvc: *httpc.NewBaseSvc(url)}
}

func (us *UserSvc) Hello(ctx context.Context, req *HelloReq) (resp *HelloResp, err error) {
    result := APIResponse{Data: &resp}
    err = us.Client().
        NewRequest().
        SetContext(ctx).
        POST(helloURL).
        SetFormData(httpc.M{"userId": req.UserId}).
        SetResult(&result).
        SendAndParse()
    return
}
```

## 3. 定义响应解析 (response.go)

实现 `httpc.IResponse` 接口：

```go
type APIResponse struct {
    Code    *int    `json:"code"`
    Message *string `json:"message"`
    Data    any     `json:"data"`
}

func (r *APIResponse) Parse(b []byte) error    { return jsonx.Unmarshal(b, &r) }
func (r *APIResponse) Valid() bool             { return r.Code != nil && r.Message != nil }
func (r *APIResponse) IsSuccess() bool         { return *r.Code == 200 }
func (r *APIResponse) Msg() string             { return *r.Message }
func (r *APIResponse) ParseData() error        { /* 解析 data 字段到目标结构 */ }
```

响应接口类型：
- `httpc.IResponse` - 标准格式（含 code/message/data）
- `httpc.IResponseNonStandard` - 非标准格式（需自定义 `ParseData([]byte)`）

## 4. 初始化 (init.go)

```go
package user

var Svc IUserSvc

func Init(url string) {
    Svc = NewUserSvc(url)
}
```

## 5. 在 Controller 中调用

```go
func (c *apiController) Index(ctx *httpx.Context) (any, error) {
    return user.Svc.Hello(ctx, &user.HelloReq{UserId: "123"})
}
```
