# Code Style Requirements

## Naming Conventions

| 类型 | 风格 | 示例 |
|-----|------|-----|
| 枚举常量 | SCREAMING_SNAKE_CASE | `USER_STATUS_NORMAL`, `ORDER_STATUS_PAID` |
| 枚举类型 | PascalCase | `UserStatus`, `OrderStatus` |
| 结构体 | PascalCase | `UserController`, `GetUsersLogic` |
| 方法/函数 | PascalCase (导出) / camelCase (私有) | `Handle()`, `parseInput()` |
| 变量 | camelCase | `userDao`, `reqData` |
| 常量 | PascalCase 或 SCREAMING_SNAKE_CASE | `MaxRetries`, `DEFAULT_TIMEOUT` |
| 包名 | 小写单词 | `httpx`, `errorx`, `logx` |
| DAO 实例 | PascalCase | `dao.User`, `dao.Order` |

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
- Logic 结构体为空结构体，直接使用 `dao.Xxx` 单例访问数据
- Handle 方法签名：`Handle(ctx context.Context, req ReqType) (*RespType, error)`

```go
type AddUserLogic struct{}

func NewAddUserLogic() *AddUserLogic {
    return &AddUserLogic{}
}

func (l *AddUserLogic) Handle(ctx context.Context, req typing.AddUserReq) (*typing.AddUserResp, error) {
    _, err := dao.User.Ctx(ctx).Data(do.User{Name: req.Name}).Insert()
    if err != nil {
        return nil, err
    }
    return &typing.AddUserResp{Message: "success"}, nil
}
```

## Model 规范

### 目录结构
```
model/
├── entity/        # 表结构定义（强类型，自动生成）
│   └── user.go
├── do/            # Data Object（any 类型，用于 Where/Data）
│   └── user.go
└── dao/           # Data Access Object
    ├── internal/  # 内部实现（自动生成）
    │   └── user.go
    └── user.go    # 外部接口（可自定义扩展）
```

### Entity 定义（自动生成）
```go
// model/entity/user.go
package entity

type User struct {
    Id        int        `gorm:"column:id;primaryKey" json:"id"`
    Name      string     `gorm:"column:name" json:"name"`
    Status    *int       `gorm:"column:status" json:"status"`
    CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
}

func (e *User) TableName() string {
    return "user"
}
```

### DO 定义（自动生成）
```go
// model/do/user.go
package do

type User struct {
    Id        any
    Name      any
    Status    any
    CreatedAt any
}
```

### DAO 定义
```go
// model/dao/user.go
package dao

type userDao struct {
    *internal.UserDao
}

var User = &userDao{internal.NewUserDao()}

// 自定义方法
func (d *userDao) GetByName(ctx context.Context, name string) (*entity.User, error) {
    var user entity.User
    err := d.Ctx(ctx).Where(do.User{Name: name}).One(&user)
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    return &user, err
}
```

## 禁止事项

1. **禁止使用 `fmt.Println`** - 使用 `logx.WithContext(ctx).Info/Debug/Error`
2. **禁止硬编码** - 配置项放入 `.env` 或 `config/`
3. **禁止忽略错误** - 必须处理或显式标注 `_ = err`
4. **禁止在 Controller 中写业务逻辑** - 业务逻辑放入 Logic 层
5. **禁止在 Logic 中实例化 DAO** - 使用 `dao.Xxx` 单例
6. **禁止在 Controller/Logic 中直接操作 db** - 使用 `dao.Xxx.Ctx(ctx)` 链式调用
7. **禁止在 Controller/Logic 中硬编码表字段名** - 使用 `do.Xxx` 或 `dao.Xxx.Columns()`

```go
// ❌ 禁止在 Logic 中直接操作 db
func (l *AddUserLogic) Handle(ctx context.Context, req typing.AddUserReq) (*typing.AddUserResp, error) {
    db.WithContext(ctx).Where("name = ?", req.Name).First(&user)  // 禁止
    db.WithContext(ctx).Model(&user).Update("status", 1)          // 禁止
}

// ✅ 正确做法：使用 DAO 链式调用
func (l *AddUserLogic) Handle(ctx context.Context, req typing.AddUserReq) (*typing.AddUserResp, error) {
    // 查询
    var user entity.User
    err := dao.User.Ctx(ctx).Where(do.User{Name: req.Name}).One(&user)

    // 更新
    _, err = dao.User.Ctx(ctx).
        Data(do.User{Status: 1}).
        Where(do.User{Id: user.Id}).
        Update()

    return &typing.AddUserResp{Message: "success"}, err
}

// ✅ 使用 Columns 常量避免硬编码
cols := dao.User.Columns()
err := dao.User.Ctx(ctx).
    Where(cols.Status+" = ?", 1).
    All(&users)
```
