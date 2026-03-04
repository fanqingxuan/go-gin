# Const 常量与枚举

常量定义在 `const/` 目录，包含两个子包：`enum/`（枚举）和 `errcode/`（业务错误码）。

## 枚举 (const/enum/)

### 定义枚举

枚举类型嵌入 `etype.BaseEnum`，常量使用 `etype.NewEnum[T]` 创建：

```go
// const/enum/user_status.go
package enum

import "go-gin/internal/etype"

type UserStatus struct {
    etype.BaseEnum
}

var (
    USER_STATUS_NORMAL   = etype.NewEnum[UserStatus](1, "正常")
    USER_STATUS_DISABLED = etype.NewEnum[UserStatus](2, "禁用")
    USER_STATUS_DELETED  = etype.NewEnum[UserStatus](3, "已删除")
)
```

```go
// const/enum/order_status.go
package enum

import "go-gin/internal/etype"

type OrderStatus struct {
    etype.BaseEnum
}

var (
    ORDER_STATUS_PENDING   = etype.NewEnum[OrderStatus](1, "待支付")
    ORDER_STATUS_PAID      = etype.NewEnum[OrderStatus](2, "已支付")
    ORDER_STATUS_SHIPPING  = etype.NewEnum[OrderStatus](3, "配送中")
    ORDER_STATUS_COMPLETED = etype.NewEnum[OrderStatus](4, "已完成")
    ORDER_STATUS_CANCELLED = etype.NewEnum[OrderStatus](5, "已取消")
)
```

### 生成代码

运行 `go run ./cmd/make/... make:enum` 自动生成 `gen_*.go` 文件，包含：
- `Scan(value any) error` - 数据库扫描
- `Value() (driver.Value, error)` - 数据库写入
- `UnmarshalJSON(data []byte) error` - JSON 反序列化
- `MarshalJSON() ([]byte, error)` - JSON 序列化

### 使用枚举

```go
// 获取 code 和描述
USER_STATUS_NORMAL.Code()    // 1
USER_STATUS_NORMAL.Desc()    // "正常"
USER_STATUS_NORMAL.String()  // "正常(1)"

// 比较
status.Equal(USER_STATUS_NORMAL)  // true/false

// 从 code 解析
parsed, err := etype.Parse[enum.UserStatus](1)  // 返回 USER_STATUS_NORMAL
```

### 命名规范

- 枚举类型：PascalCase，如 `UserStatus`、`OrderStatus`
- 枚举常量：SCREAMING_SNAKE_CASE，前缀为类型名，如 `USER_STATUS_NORMAL`、`ORDER_STATUS_PAID`

## 业务错误码 (const/errcode/)

### 定义错误码

错误码从 20000 开始（10000-20000 保留给框架）：

```go
// const/errcode/user.go
package errcode

import "go-gin/internal/errorx"

var (
    ErrUserNotFound       = errorx.New(20001, "用户不存在")
    ErrUserNameOrPwdFaild = errorx.New(20002, "用户名或者密码错误")
    ErrUserMustLogin      = errorx.New(20003, "请先登录")
    ErrUserNeedLoginAgain = errorx.New(20004, "token已过期,请重新登录")
)
```

### 工具函数

```go
// const/errcode/util.go
// 快速创建默认业务错误（code 为 ErrCodeBizDefault = 10001）
errcode.NewDefault("自定义错误信息")
```

### 在 Logic 中使用

```go
func (l *LoginLogic) Handle(ctx context.Context, req typing.LoginReq) (*typing.LoginResp, error) {
    var user entity.User
    err := dao.User.Ctx(ctx).Where(do.User{Name: req.Username}).One(&user)
    if err != nil {
        return nil, err
    }
    if user.Id == 0 {
        return nil, errcode.ErrUserNotFound
    }
    // ...
}
```

### 框架预留错误码

| 错误码 | 说明 |
|--------|------|
| 10000 | 默认通用错误 |
| 10001 | 业务默认错误 |
| 10002 | 验证失败 |
| 10003 | 数据库操作失败 |
| 10004 | Redis 操作失败 |
| 11001-11006 | 第三方接口错误 |
