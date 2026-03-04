# Logic 业务逻辑层

`logic/` 目录负责业务逻辑处理，采用 Command 模式，每个操作一个 Logic 结构体。

## 基本结构

每个 Logic 文件包含：空结构体、构造函数、Handle 方法。

```go
// logic/add_user_logic.go
package logic

import (
    "context"
    "go-gin/model/dao"
    "go-gin/model/do"
    "go-gin/typing"
)

type AddUserLogic struct{}

func NewAddUserLogic() *AddUserLogic {
    return &AddUserLogic{}
}

func (l *AddUserLogic) Handle(ctx context.Context, req typing.AddUserReq) (*typing.AddUserResp, error) {
    id, err := dao.User.Ctx(ctx).Data(do.User{Name: req.Name}).InsertAndGetId()
    if err != nil {
        return nil, err
    }
    return &typing.AddUserResp{
        Message: fmt.Sprintf("message:%d", id),
    }, nil
}
```

## LogicHandler 接口

Logic 必须实现 `httpx.LogicHandler` 泛型接口：

```go
type LogicHandler[Req any, Resp any] interface {
    Handle(ctx context.Context, req Req) (Resp, error)
}
```

## Handle 签名模式

### 标准模式（typed request + typed response）

```go
func (l *AddUserLogic) Handle(ctx context.Context, req typing.AddUserReq) (*typing.AddUserResp, error)
```

### 无请求参数模式

请求类型使用 `any`，适用于无需参数的操作：

```go
func (l *IndexLogic) Handle(ctx context.Context, req any) (any, error) {
    return "user/index", nil
}
```

### 查询列表模式

结合 DAO 查询和 Transformer 转换：

```go
func (l *GetUsersLogic) Handle(ctx context.Context, req typing.ListReq) (*typing.ListResp, error) {
    var users []*entity.User
    err := dao.User.Ctx(ctx).All(&users)
    if errorx.IsError(err) {
        return nil, err
    }
    return &typing.ListResp{
        Data: transformer.ConvertUserToListData(users),
    }, nil
}
```

### 批量写入模式

```go
func (l *MultiAddUserLogic) Handle(ctx context.Context, req typing.MultiUserAddReq) (*typing.MultiUserAddResp, error) {
    users := make([]do.User, len(req.Users))
    for i, user := range req.Users {
        users[i] = do.User{Name: user.Name}
    }
    _, err := dao.User.Ctx(ctx).Data(users).Insert()
    if err != nil {
        return nil, err
    }
    return &typing.MultiUserAddResp{Message: "批量添加成功"}, nil
}
```

### 业务错误返回模式

使用 `errcode` 返回业务错误：

```go
func (l *LoginLogic) Handle(ctx context.Context, req typing.LoginReq) (*typing.LoginResp, error) {
    user, err := dao.User.GetByName(ctx, req.Username)
    if err != nil {
        return nil, err
    }
    if user == nil {
        return nil, errcode.ErrUserNotFound
    }
    // ...
    return &typing.LoginResp{Token: t, User: *user}, nil
}
```

## 与 Controller 的关联

Controller 通过 `httpx.Handle` 自动绑定请求并调用 Logic：

```go
// controller
func (c *userController) AddUser(ctx *httpx.Context) (any, error) {
    return httpx.Handle(ctx, logic.NewAddUserLogic())
}
```

框架自动完成：参数绑定 → 验证 → 调用 `Handle` → 格式化响应。

## 命名规范

| 项目 | 规范 | 示例 |
|------|------|------|
| 文件名 | `{action}_logic.go` | `add_user_logic.go`、`get_users_logic.go` |
| 结构体 | `{Action}Logic` | `AddUserLogic`、`GetUsersLogic` |
| 构造函数 | `New{Action}Logic()` | `NewAddUserLogic()` |
| 方法名 | 固定为 `Handle` | `Handle(ctx, req)` |

## 禁止事项

1. **禁止在 Logic 中实例化 DAO** —— 使用 `dao.Xxx` 单例
2. **禁止在 Logic 中直接操作 db** —— 使用 `dao.Xxx.Ctx(ctx)` 链式调用
3. **禁止在 Logic 中硬编码字段名** —— 使用 `do.Xxx` 或 `dao.Xxx.Columns()`
4. **禁止在 Logic 中使用 `fmt.Println`** —— 使用 `logx.WithContext(ctx)`
