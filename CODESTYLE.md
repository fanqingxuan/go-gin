# Code Style Requirements

## Naming Conventions

| 类型 | 风格 | 示例 |
|-----|------|-----|
| 枚举常量 | SCREAMING_SNAKE_CASE | `USER_STATUS_NORMAL`, `ORDER_STATUS_PAID` |
| 枚举类型 | PascalCase | `UserStatus`, `OrderStatus` |
| 结构体 | PascalCase | `UserController`, `GetUsersLogic` |
| 方法/函数 | PascalCase (导出) / camelCase (私有) | `Handle()`, `parseInput()` |
| 变量 | camelCase | `userModel`, `reqData` |
| 常量 | PascalCase 或 SCREAMING_SNAKE_CASE | `MaxRetries`, `DEFAULT_TIMEOUT` |
| 包名 | 小写单词 | `httpx`, `errorx`, `logx` |

## Controller 规范

统一使用 `httpx.ShouldBindHandle` 方式调用 Logic：
```go
// ✅ 推荐
func (c *userController) AddUser(ctx *httpx.Context) (any, error) {
    return httpx.ShouldBindHandle(ctx, logic.NewAddUserLogic())
}

// ❌ 避免手动绑定
func (c *userController) AddUser(ctx *httpx.Context) (any, error) {
    var req typing.AddUserReq
    if err := ctx.ShouldBind(&req); err != nil { ... }
    return logic.NewAddUserLogic().Handle(ctx, req)
}
```

## Logic 规范

- 每个 Logic 一个文件，文件名格式：`{action}_logic.go`
- Logic 结构体包含依赖的 model
- Handle 方法签名：`Handle(ctx context.Context, req ReqType) (*RespType, error)`

```go
type AddUserLogic struct {
    model *model.UserModel
}

func NewAddUserLogic() *AddUserLogic {
    return &AddUserLogic{
        model: model.NewUserModel(),
    }
}

func (l *AddUserLogic) Handle(ctx context.Context, req typing.AddUserReq) (*typing.AddUserResp, error) {
    // 业务逻辑
}
```

## 禁止事项

1. **禁止使用 `fmt.Println`** - 使用 `logx.WithContext(ctx).Info/Debug/Error`
2. **禁止硬编码** - 配置项放入 `.env` 或 `config/`
3. **禁止忽略错误** - 必须处理或显式标注 `_ = err`
4. **禁止在 Controller 中写业务逻辑** - 业务逻辑放入 Logic 层
