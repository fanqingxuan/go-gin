# DAO 链式调用

基于 GORM 封装的 GoFrame 风格 Model 链式 API，定义在 `internal/component/db/model.go`。

## 基本用法

所有数据库操作必须通过 `dao.Xxx.Ctx(ctx)` 开始链式调用。

## 查询

```go
// 查询单条
var user entity.User
err := dao.User.Ctx(ctx).Where(do.User{Id: 1}).One(&user)

// 查询单条 + 判断是否存在
found, err := dao.User.Ctx(ctx).Where(do.User{Name: "test"}).Found(&user)

// 查询列表
var users []entity.User
err := dao.User.Ctx(ctx).Where(do.User{Status: 1}).All(&users)

// 分页查询
err := dao.User.Ctx(ctx).
    Where(do.User{Status: 1}).
    Order("id DESC").
    Page(1, 10).
    All(&users)

// 分页查询 + 总数
var total int64
err := dao.User.Ctx(ctx).
    Where(do.User{Status: 1}).
    Page(page, size).
    ScanAndCount(&users, &total)

// 指定字段
cols := dao.User.Columns()
err := dao.User.Ctx(ctx).
    Fields(cols.Id, cols.Name).
    Where(cols.Status+" = ?", 1).
    All(&users)

// 排除字段
err := dao.User.Ctx(ctx).FieldsEx("password").All(&users)
```

## 写入

```go
// 插入
_, err := dao.User.Ctx(ctx).Data(do.User{Name: "test", Status: 1}).Insert()

// 插入并获取 ID
id, err := dao.User.Ctx(ctx).Data(do.User{Name: "test"}).InsertAndGetId()

// INSERT IGNORE
_, err := dao.User.Ctx(ctx).Data(do.User{Name: "test"}).InsertIgnore()

// REPLACE INTO
_, err := dao.User.Ctx(ctx).Data(do.User{Name: "test"}).Replace()

// 更新（必须有 Where 条件）
_, err := dao.User.Ctx(ctx).
    Data(do.User{Status: 2}).
    Where(do.User{Id: 1}).
    Update()

// 删除（必须有 Where 条件）
_, err := dao.User.Ctx(ctx).Where(do.User{Id: 1}).Delete()

// Save (INSERT ... ON DUPLICATE KEY UPDATE)
_, err := dao.User.Ctx(ctx).Data(userData).Save()
```

## 条件构建器

```go
Where(do.User{Status: 1})          // 结构体条件
Where("age > ?", 18)               // 原生条件
WhereOr(do.User{Status: 2})        // OR 条件
WherePri(1)                        // 主键条件
WhereLT("age", 18)                 // age < 18
WhereLTE("age", 18)                // age <= 18
WhereGT("age", 18)                 // age > 18
WhereGTE("age", 18)                // age >= 18
WhereBetween("age", 18, 30)        // BETWEEN 18 AND 30
WhereNotBetween("age", 18, 30)     // NOT BETWEEN
WhereLike("name", "%test%")        // LIKE
WhereNotLike("name", "%test%")     // NOT LIKE
WhereIn("id", []int{1, 2, 3})     // IN
WhereNotIn("id", []int{1, 2})     // NOT IN
WhereNull("deleted_at")            // IS NULL
WhereNotNull("deleted_at")         // IS NOT NULL
WhereNot("status", 0)              // status != 0
Wheref("FIND_IN_SET(%d, tags)", 5) // 格式化（仅用于动态列名，禁止用户输入）
```

## 聚合

```go
count, err := dao.User.Ctx(ctx).Where(do.User{Status: 1}).Count()
exists, err := dao.User.Ctx(ctx).Where(do.User{Id: 1}).Exist()
min, err := dao.Order.Ctx(ctx).Min("amount")
max, err := dao.Order.Ctx(ctx).Max("amount")
avg, err := dao.Order.Ctx(ctx).Avg("amount")
sum, err := dao.Order.Ctx(ctx).Sum("amount")
```

## 字段提取

```go
// 获取单列切片
var names []string
err := dao.User.Ctx(ctx).Pluck("name", &names)

// 获取单个值
var name string
err := dao.User.Ctx(ctx).Where(do.User{Id: 1}).Value("name", &name)

// 获取列值为 []any
arr, err := dao.User.Ctx(ctx).Array("name")
```

## 递增/递减

```go
dao.User.Ctx(ctx).Where(do.User{Id: 1}).Increment("score", 10)
dao.User.Ctx(ctx).Where(do.User{Id: 1}).Decrement("balance", 100)
```

## 分块处理

```go
dao.User.Ctx(ctx).Where(do.User{Status: 1}).Chunk(100, func(result []map[string]any, err error) bool {
    // 处理每批数据
    return true // 返回 false 停止
})
```

## 链式方法

| 方法 | 说明 |
|------|------|
| `Fields(fields...)` | 指定查询字段 |
| `FieldsEx(fields...)` | 排除字段 |
| `Order(order)` | 排序 |
| `Group(group)` | 分组 |
| `Having(having)` | 分组过滤 |
| `Limit(n)` | 限制条数 |
| `Offset(n)` | 偏移量 |
| `Page(page, size)` | 分页 |
| `Distinct()` | 去重 |
| `Unscoped()` | 忽略软删除 |
| `Data(data)` | 设置写入数据 |
| `SetPrimaryKey(pk)` | 设置主键字段 |
