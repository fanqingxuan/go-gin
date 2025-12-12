// 测试 db 和 model 层操作
// 用法: go run scripts/test_db/main.go -f .env
package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go-gin/internal/component/db"
	"go-gin/internal/errorx"
	"go-gin/model/dao"
	"go-gin/model/do"
	"go-gin/model/entity"
	"go-gin/scripts"
)

func main() {
	scripts.Init()
	ctx := context.Background()

	fmt.Println("=== 测试 db 包 ===")
	testDBPackage(ctx)

	fmt.Println("\n=== 测试 Model 链式调用 ===")
	testModelChain(ctx)

	fmt.Println("\n=== 测试 DAO 层 ===")
	testDAO(ctx)

	fmt.Println("\n=== 测试 DBError 类型 ===")
	testDBError(ctx)

	fmt.Println("\n=== 所有测试完成 ===")
}

func testDBPackage(ctx context.Context) {
	// 测试 Ping
	if err := db.Ping(ctx); err != nil {
		log.Fatalf("Ping 失败: %v", err)
	}
	fmt.Println("✓ Ping 成功")

	// 测试 Query
	result, err := db.Query(ctx, "SELECT 1 as num")
	if err != nil {
		log.Fatalf("Query 失败: %v", err)
	}
	fmt.Printf("✓ Query 成功: %v\n", result)

	// 测试 GetOne
	one, err := db.GetOne(ctx, "SELECT 1 as num, 'test' as name")
	if err != nil {
		log.Fatalf("GetOne 失败: %v", err)
	}
	fmt.Printf("✓ GetOne 成功: %v\n", one)

	// 测试 GetValue
	val, err := db.GetValue(ctx, "SELECT COUNT(*) FROM user")
	if err != nil {
		log.Fatalf("GetValue 失败: %v", err)
	}
	fmt.Printf("✓ GetValue 成功: %v\n", val)

	// 测试 WithContext + Raw
	var count int64
	err = db.WithContext(ctx).Raw("SELECT COUNT(*) FROM user").Scan(&count).Error()
	if err != nil {
		log.Fatalf("WithContext.Raw 失败: %v", err)
	}
	fmt.Printf("✓ WithContext.Raw 成功: count=%d\n", count)
}

func testModelChain(ctx context.Context) {
	// 测试 Count
	count, err := db.NewModel(ctx, "user").Count()
	if err != nil {
		log.Fatalf("Count 失败: %v", err)
	}
	fmt.Printf("✓ Count 成功: %d\n", count)

	// 测试 Where + All
	var users []map[string]any
	err = db.NewModel(ctx, "user").
		Where("status = ?", 1).
		Order("id DESC").
		Limit(5).
		Scan(&users)
	if err != nil {
		var dbErr errorx.DBError
		if errors.As(err, &dbErr) {
			fmt.Printf("✓ DBError 类型检查: %s\n", dbErr.Msg)
		} else {
			log.Fatalf("Where+All 失败: %v", err)
		}
	}
	fmt.Printf("✓ Where+All 成功: 查询到 %d 条记录\n", len(users))

	// 测试 Page
	var pageUsers []map[string]any
	err = db.NewModel(ctx, "user").
		Page(1, 10).
		Scan(&pageUsers)
	if err != nil {
		log.Fatalf("Page 失败: %v", err)
	}
	fmt.Printf("✓ Page 成功: 查询到 %d 条记录\n", len(pageUsers))

	// 测试 Exist
	exist, err := db.NewModel(ctx, "user").Where("id = ?", 1).Exist()
	if err != nil {
		log.Fatalf("Exist 失败: %v", err)
	}
	fmt.Printf("✓ Exist 成功: %v\n", exist)

	// 测试 ScanAndCount
	var allUsers []entity.User
	var total int64
	err = db.NewModel(ctx, "user").
		Page(1, 5).
		ScanAndCount(&allUsers, &total)
	if err != nil {
		log.Fatalf("ScanAndCount 失败: %v", err)
	}
	fmt.Printf("✓ ScanAndCount 成功: 本页 %d 条, 总共 %d 条\n", len(allUsers), total)

	// 测试 Where 变体
	testWhereVariants(ctx)

	// 测试聚合函数
	testAggregateFunctions(ctx)

	// 测试其他方法
	testOtherMethods(ctx)
}

func testWhereVariants(ctx context.Context) {
	fmt.Println("\n--- Where 变体测试 ---")

	// WhereLT
	c, _ := db.NewModel(ctx, "user").WhereLT("id", 100).Count()
	fmt.Printf("✓ WhereLT: id < 100 有 %d 条\n", c)

	// WhereLTE
	c, _ = db.NewModel(ctx, "user").WhereLTE("id", 100).Count()
	fmt.Printf("✓ WhereLTE: id <= 100 有 %d 条\n", c)

	// WhereGT
	c, _ = db.NewModel(ctx, "user").WhereGT("id", 0).Count()
	fmt.Printf("✓ WhereGT: id > 0 有 %d 条\n", c)

	// WhereGTE
	c, _ = db.NewModel(ctx, "user").WhereGTE("id", 1).Count()
	fmt.Printf("✓ WhereGTE: id >= 1 有 %d 条\n", c)

	// WhereBetween
	c, _ = db.NewModel(ctx, "user").WhereBetween("id", 1, 100).Count()
	fmt.Printf("✓ WhereBetween: id BETWEEN 1 AND 100 有 %d 条\n", c)

	// WhereLike
	c, _ = db.NewModel(ctx, "user").WhereLike("name", "%test%").Count()
	fmt.Printf("✓ WhereLike: name LIKE '%%test%%' 有 %d 条\n", c)

	// WhereIn
	c, _ = db.NewModel(ctx, "user").WhereIn("id", []int{1, 2, 3}).Count()
	fmt.Printf("✓ WhereIn: id IN (1,2,3) 有 %d 条\n", c)

	// WhereNotIn
	c, _ = db.NewModel(ctx, "user").WhereNotIn("id", []int{999}).Count()
	fmt.Printf("✓ WhereNotIn: id NOT IN (999) 有 %d 条\n", c)

	// WhereNull
	c, _ = db.NewModel(ctx, "user").WhereNull("deleted_at").Count()
	fmt.Printf("✓ WhereNull: deleted_at IS NULL 有 %d 条\n", c)

	// WhereNotNull
	c, _ = db.NewModel(ctx, "user").WhereNotNull("name").Count()
	fmt.Printf("✓ WhereNotNull: name IS NOT NULL 有 %d 条\n", c)

	// WhereNot
	c, _ = db.NewModel(ctx, "user").WhereNot("status", 999).Count()
	fmt.Printf("✓ WhereNot: status != 999 有 %d 条\n", c)

	// Wheref
	c, _ = db.NewModel(ctx, "user").Wheref("id > %d", 0).Count()
	fmt.Printf("✓ Wheref: id > 0 有 %d 条\n", c)

	// WhereOr
	c, _ = db.NewModel(ctx, "user").Where("id = ?", 1).WhereOr("id = ?", 2).Count()
	fmt.Printf("✓ WhereOr: id=1 OR id=2 有 %d 条\n", c)
}

func testAggregateFunctions(ctx context.Context) {
	fmt.Println("\n--- 聚合函数测试 ---")

	// Min
	min, err := db.NewModel(ctx, "user").Min("id")
	if err != nil {
		log.Fatalf("Min 失败: %v", err)
	}
	fmt.Printf("✓ Min(id): %.0f\n", min)

	// Max
	max, err := db.NewModel(ctx, "user").Max("id")
	if err != nil {
		log.Fatalf("Max 失败: %v", err)
	}
	fmt.Printf("✓ Max(id): %.0f\n", max)

	// Avg
	avg, err := db.NewModel(ctx, "user").Avg("id")
	if err != nil {
		log.Fatalf("Avg 失败: %v", err)
	}
	fmt.Printf("✓ Avg(id): %.2f\n", avg)

	// Sum
	sum, err := db.NewModel(ctx, "user").Sum("id")
	if err != nil {
		log.Fatalf("Sum 失败: %v", err)
	}
	fmt.Printf("✓ Sum(id): %.0f\n", sum)
}

func testOtherMethods(ctx context.Context) {
	fmt.Println("\n--- 其他方法测试 ---")

	// Fields
	var users []map[string]any
	err := db.NewModel(ctx, "user").Fields("id", "name").Limit(3).Scan(&users)
	if err != nil {
		log.Fatalf("Fields 失败: %v", err)
	}
	fmt.Printf("✓ Fields(id, name): %v\n", users)

	// Distinct
	c, _ := db.NewModel(ctx, "user").Distinct().Fields("status").Count()
	fmt.Printf("✓ Distinct status count: %d\n", c)

	// Array
	arr, err := db.NewModel(ctx, "user").Limit(5).Array("id")
	if err != nil {
		log.Fatalf("Array 失败: %v", err)
	}
	fmt.Printf("✓ Array(id): %v\n", arr)

	// Pluck
	var ids []int64
	err = db.NewModel(ctx, "user").Limit(5).Pluck("id", &ids)
	if err != nil {
		log.Fatalf("Pluck 失败: %v", err)
	}
	fmt.Printf("✓ Pluck(id): %v\n", ids)

	// Value
	var name string
	err = db.NewModel(ctx, "user").Where("id > ?", 0).Value("name", &name)
	if err != nil {
		log.Fatalf("Value 失败: %v", err)
	}
	fmt.Printf("✓ Value(name): %s\n", name)

	// Group + Having
	var grouped []map[string]any
	err = db.NewModel(ctx, "user").
		Fields("status", "COUNT(*) as cnt").
		Group("status").
		Having("cnt > 0").
		Scan(&grouped)
	if err != nil {
		log.Fatalf("Group+Having 失败: %v", err)
	}
	fmt.Printf("✓ Group+Having: %v\n", grouped)

	// Chunk
	fmt.Print("✓ Chunk: ")
	chunkCount := 0
	db.NewModel(ctx, "user").Chunk(2, func(result []map[string]any, err error) bool {
		if err != nil {
			return false
		}
		chunkCount += len(result)
		fmt.Printf("[批次%d条] ", len(result))
		return true
	})
	fmt.Printf("共 %d 条\n", chunkCount)
}

func testDAO(ctx context.Context) {
	// 测试 DAO Count
	count, err := dao.User.Ctx(ctx).Count()
	if err != nil {
		log.Fatalf("DAO Count 失败: %v", err)
	}
	fmt.Printf("✓ DAO Count 成功: %d\n", count)

	// 测试 DAO Where + One
	var user entity.User
	err = dao.User.Ctx(ctx).
		Where(do.User{Status: 1}).
		Order("id ASC").
		One(&user)
	if err != nil {
		log.Fatalf("DAO Where+One 失败: %v", err)
	}
	if user.Id > 0 {
		fmt.Printf("✓ DAO Where+One 成功: id=%d, name=%s\n", user.Id, user.Name)
	} else {
		fmt.Println("✓ DAO Where+One 成功: 无记录")
	}

	// 测试 DAO Columns
	cols := dao.User.Columns()
	fmt.Printf("✓ DAO Columns 成功: Id=%s, Name=%s, Status=%s\n", cols.Id, cols.Name, cols.Status)

	// 测试 DAO WherePri
	var userById entity.User
	err = dao.User.Ctx(ctx).WherePri(1).One(&userById)
	if err != nil {
		log.Fatalf("DAO WherePri 失败: %v", err)
	}
	if userById.Id > 0 {
		fmt.Printf("✓ DAO WherePri 成功: id=%d\n", userById.Id)
	} else {
		fmt.Println("✓ DAO WherePri 成功: 无记录")
	}

	// 测试 DAO All
	var allUsers []entity.User
	err = dao.User.Ctx(ctx).
		Where("status = ?", 1).
		Limit(3).
		All(&allUsers)
	if err != nil {
		log.Fatalf("DAO All 失败: %v", err)
	}
	fmt.Printf("✓ DAO All 成功: 查询到 %d 条记录\n", len(allUsers))

	// 测试事务（只读，不实际修改数据）
	err = db.Transaction(ctx, func(tx *db.TX) error {
		var txUser entity.User
		err := tx.Model("user").Where("id = ?", 1).One(&txUser)
		if err != nil {
			return err
		}
		fmt.Printf("✓ Transaction 查询成功: id=%d\n", txUser.Id)
		return nil
	})
	if err != nil {
		log.Fatalf("Transaction 失败: %v", err)
	}
	fmt.Println("✓ Transaction 成功")
}

func testDBError(ctx context.Context) {
	// 故意使用错误的字段名触发数据库错误
	_, err := db.NewModel(ctx, "user").Where("invalid_column = ?", 1).Count()
	if err != nil {
		var dbErr errorx.DBError
		if errors.As(err, &dbErr) {
			fmt.Printf("✓ errors.As(err, &DBError) 成功: %s\n", dbErr.Msg)
		} else {
			log.Fatalf("DBError 类型检查失败: %T", err)
		}
	}

	// 测试 ErrRecordNotFound 不会被包装
	var user entity.User
	err = db.NewModel(ctx, "user").Where("id = ?", -1).One(&user)
	if err == nil {
		fmt.Println("✓ ErrRecordNotFound 返回 nil (符合预期)")
	}
}
