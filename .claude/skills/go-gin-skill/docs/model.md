# Model 数据模型

`model/` 目录包含三个子包，由 `make:dao` 自动生成，形成类型安全的数据访问层。

## 目录结构

```
model/
├── entity/        # 表结构（强类型，GORM 模型）
│   └── user.go
├── do/            # Data Object（any 类型，用于 Where/Data 条件）
│   └── user.go
└── dao/           # Data Access Object（单例 + 链式 API）
    ├── internal/  # 内部实现（自动生成，勿修改）
    │   └── user.go
    └── user.go    # 外部接口（可自定义扩展）
```

## Entity（表结构）

自动生成，强类型映射数据库表，可空字段使用指针类型：

```go
// model/entity/user.go（自动生成，勿修改）
package entity

import "time"

type User struct {
    Id        int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
    Name      string     `gorm:"column:name" json:"name"`
    Status    *int       `gorm:"column:status" json:"status"`
    CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
    UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
    DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
    Age       *int       `gorm:"column:age" json:"age"`
}

func (e *User) TableName() string {
    return "user"
}
```

**规范**：
- 主键使用 `int`，非空字段使用值类型，可空字段使用指针（`*int`、`*time.Time`）
- GORM tag 定义列名、主键、自增等约束
- JSON tag 用于序列化

## DO（Data Object）

自动生成，所有字段为 `any` 类型，用于 `Where()` 和 `Data()` 条件构建，仅非 nil 字段参与查询：

```go
// model/do/user.go（自动生成，勿修改）
package do

type User struct {
    Id        any
    Name      any
    Status    any
    CreatedAt any
    UpdatedAt any
    DeletedAt any
    Age       any
}
```

**使用场景**：
```go
// Where 条件 —— 仅 Status 参与查询
dao.User.Ctx(ctx).Where(do.User{Status: 1}).All(&users)

// Data 写入 —— 仅 Name、Status 参与插入
dao.User.Ctx(ctx).Data(do.User{Name: "test", Status: 1}).Insert()
```

## DAO（Data Access Object）

### Internal（自动生成）

提供基础的 `Ctx()`、`Columns()`、`Transaction()` 方法：

```go
// model/dao/internal/user.go（自动生成，勿修改）
package internal

type UserDao struct {
    table      string
    primaryKey string
    columns    UserColumns
}

type UserColumns struct {
    Id        string
    Name      string
    Status    string
    CreatedAt string
    UpdatedAt string
    DeletedAt string
    Age       string
}

func NewUserDao() *UserDao {
    return &UserDao{
        table:      "user",
        primaryKey: "id",
        columns:    userColumns,
    }
}

func (d *UserDao) Ctx(ctx context.Context) *db.Model {
    return db.NewModel(ctx, d.table).SetPrimaryKey(d.primaryKey)
}

func (d *UserDao) Columns() UserColumns {
    return d.columns
}

func (d *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *db.TX) error) error {
    return db.Transaction(ctx, func(tx *db.TX) error {
        return f(ctx, tx)
    })
}
```

### Public（可自定义扩展）

嵌入 internal DAO，暴露为包级单例，可添加自定义方法：

```go
// model/dao/user.go
package dao

import (
    "context"
    "go-gin/model/dao/internal"
    "go-gin/model/do"
    "go-gin/model/entity"
)

type userDao struct {
    *internal.UserDao
}

var User = &userDao{internal.NewUserDao()}

// 自定义方法
func (d *userDao) GetByName(ctx context.Context, name string) (*entity.User, error) {
    var user entity.User
    found, err := d.Ctx(ctx).Where(do.User{Name: name}).Found(&user)
    if err != nil {
        return nil, err
    }
    if !found {
        return nil, nil
    }
    return &user, nil
}
```

## 代码生成

```bash
# 生成所有表的 entity/do/dao
go run ./cmd/make/... make:dao -f .env

# 生成指定表
go run ./cmd/make/... make:dao -f .env -t user
```

生成规则：
- `entity/` 和 `do/` 文件每次**覆盖重写**
- `dao/internal/` 文件每次**覆盖重写**
- `dao/*.go`（public）仅首次生成，**不会覆盖**，可安全添加自定义方法

## 使用规范

| 规则 | 说明 |
|------|------|
| 必须通过 `dao.Xxx.Ctx(ctx)` 访问 | 禁止在 Logic 中直接操作 db |
| 使用 `do.Xxx` 构建条件 | 禁止硬编码字段名字符串 |
| 使用 `entity.Xxx` 接收结果 | 强类型保证 |
| 使用 `dao.Xxx.Columns()` | 需要字段名常量时使用 |
| 自定义方法写在 public DAO | 禁止修改 internal 文件 |
| 单例访问 | 禁止在 Logic 中实例化 DAO |
