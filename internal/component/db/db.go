package db

import (
	"go-gin/internal/errorx"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DB struct {
	*gorm.DB
}

// 查询方法
func (db *DB) Where(query any, args ...any) *DB {
	return &DB{db.DB.Where(query, args...)}
}

func (db *DB) Select(query any, args ...any) *DB {
	return &DB{db.DB.Select(query, args...)}
}

func (db *DB) First(dest any, conds ...any) *DB {
	return &DB{db.DB.First(dest, conds...)}
}

func (db *DB) Last(dest any, conds ...any) *DB {
	return &DB{db.DB.Last(dest, conds...)}
}

func (db *DB) Find(dest any, conds ...any) *DB {
	return &DB{db.DB.Find(dest, conds...)}
}

func (db *DB) Take(dest any, conds ...any) *DB {
	return &DB{db.DB.Take(dest, conds...)}
}

func (db *DB) Scan(dest any) *DB {
	return &DB{db.DB.Scan(dest)}
}

// 创建方法
func (db *DB) Create(value any) *DB {
	return &DB{db.DB.Create(value)}
}

func (db *DB) CreateInBatches(value any, batchSize int) *DB {
	return &DB{db.DB.CreateInBatches(value, batchSize)}
}

// 更新方法
func (db *DB) Save(value any) *DB {
	return &DB{db.DB.Save(value)}
}

func (db *DB) Updates(values any) *DB {
	return &DB{db.DB.Updates(values)}
}

func (db *DB) Update(column string, value any) *DB {
	return &DB{db.DB.Update(column, value)}
}

// 删除方法
func (db *DB) Delete(value any, conds ...any) *DB {
	return &DB{db.DB.Delete(value, conds...)}
}

// 条件构造方法
func (db *DB) Or(query any, args ...any) *DB {
	return &DB{db.DB.Or(query, args...)}
}

func (db *DB) Not(query any, args ...any) *DB {
	return &DB{db.DB.Not(query, args...)}
}

func (db *DB) Distinct(args ...any) *DB {
	return &DB{db.DB.Distinct(args...)}
}

func (db *DB) Omit(columns ...string) *DB {
	return &DB{db.DB.Omit(columns...)}
}

// 分页和排序
func (db *DB) Limit(limit int) *DB {
	return &DB{db.DB.Limit(limit)}
}

func (db *DB) Offset(offset int) *DB {
	return &DB{db.DB.Offset(offset)}
}

func (db *DB) Order(value any) *DB {
	return &DB{db.DB.Order(value)}
}

func (db *DB) Group(name string) *DB {
	return &DB{db.DB.Group(name)}
}

func (db *DB) Having(query any, args ...any) *DB {
	return &DB{db.DB.Having(query, args...)}
}

// 关联查询
func (db *DB) Joins(query string, args ...any) *DB {
	return &DB{db.DB.Joins(query, args...)}
}

func (db *DB) Preload(query string, args ...any) *DB {
	return &DB{db.DB.Preload(query, args...)}
}

// 事务相关
func (db *DB) Begin() *DB {
	return &DB{db.DB.Begin()}
}

func (db *DB) Commit() *DB {
	return &DB{db.DB.Commit()}
}

func (db *DB) Rollback() *DB {
	return &DB{db.DB.Rollback()}
}

// 锁相关
func (db *DB) Clauses(conds ...clause.Expression) *DB {
	return &DB{db.DB.Clauses(conds...)}
}

// 统计
func (db *DB) Count(count *int64) *DB {
	return &DB{db.DB.Count(count)}
}

// 原始SQL
func (db *DB) Raw(sql string, values ...any) *DB {
	return &DB{db.DB.Raw(sql, values...)}
}

func (db *DB) Exec(sql string, values ...any) *DB {
	return &DB{db.DB.Exec(sql, values...)}
}

// Scopes
func (db *DB) Scopes(funcs ...func(*gorm.DB) *gorm.DB) *DB {
	return &DB{db.DB.Scopes(funcs...)}
}

// 错误处理
func (db *DB) Error() error {
	return errorx.TryToDBError(db.DB.Error)
}

// Ping
func (db *DB) Ping() error {
	var num int
	err := db.DB.Raw("select 1").Scan(&num).Error
	return errorx.TryToDBError(err)
}
