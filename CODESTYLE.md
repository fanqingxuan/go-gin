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
    user := entity.User{Name: req.Name}
    if err := dao.User.Create(ctx, &user); err != nil {
        return nil, err
    }
    return &typing.AddUserResp{Message: "success"}, nil
}
```

## Model 规范

### 目录结构
```
model/
├── dao/           # Data Access Object
│   ├── base.go    # BaseDao 通用 CRUD 方法
│   └── user.go    # 具体 DAO 实现
└── entity/        # 表结构定义
    └── user.go    # User 实体 + BaseEntity
```

### Entity 定义
```go
// model/entity/user.go
package entity

type User struct {
    BaseEntity
    Id   int64
    Name string `gorm:"column:name" json:"name"`
}

func (u *User) TableName() string {
    return "user"
}
```

### DAO 定义
```go
// model/dao/user.go
package dao

type userDao struct {
    BaseDao[entity.User]
}

var User = &userDao{}  // 单例，直接使用 dao.User

// 自定义方法
func (d *userDao) GetByName(ctx context.Context, name string) (*entity.User, error) {
    // ...
}
```

## 禁止事项

1. **禁止使用 `fmt.Println`** - 使用 `logx.WithContext(ctx).Info/Debug/Error`
2. **禁止硬编码** - 配置项放入 `.env` 或 `config/`
3. **禁止忽略错误** - 必须处理或显式标注 `_ = err`
4. **禁止在 Controller 中写业务逻辑** - 业务逻辑放入 Logic 层
5. **禁止在 Logic 中实例化 DAO** - 使用 `dao.Xxx` 单例
6. **禁止在 Controller/Logic 中直接操作 db** - 数据库操作必须封装在 DAO 层
7. **禁止在 Controller/Logic 中硬编码表字段名** - 字段操作封装在 DAO 层

```go
// ❌ 禁止在 Logic 中直接操作 db
func (l *AddUserLogic) Handle(ctx context.Context, req typing.AddUserReq) (*typing.AddUserResp, error) {
    db.WithContext(ctx).Where("name = ?", req.Name).First(&user)  // 禁止
    db.WithContext(ctx).Model(&user).Update("status", 1)          // 禁止
}

// ✅ 正确做法：在 DAO 中封装
// dao/user.go
func (d *userDao) GetByName(ctx context.Context, name string) (*entity.User, error) {
    var user entity.User
    result := db.WithContext(ctx).First(&user, "name = ?", name)
    return &user, result.Error()
}

func (d *userDao) UpdateStatus(ctx context.Context, id int64, status int) error {
    return db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", id).Update("status", status).Error()
}

// logic/xxx_logic.go
func (l *AddUserLogic) Handle(ctx context.Context, req typing.AddUserReq) (*typing.AddUserResp, error) {
    user, err := dao.User.GetByName(ctx, req.Name)  // 正确
    dao.User.UpdateStatus(ctx, user.Id, 1)          // 正确
}
```
