package db

import (
	"context"
	"errors"
	"go-gin/internal/errorx"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DB struct {
	*gorm.DB
}

// 上下文
func (db *DB) WithContext(ctx context.Context) *DB {
	return &DB{db.DB.WithContext(ctx)}
}

// 模型/表
func (db *DB) Model(value any) *DB {
	return &DB{db.DB.Model(value)}
}

func (db *DB) Table(name string, args ...any) *DB {
	return &DB{db.DB.Table(name, args...)}
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

func (db *DB) Pluck(column string, dest any) *DB {
	return &DB{db.DB.Pluck(column, dest)}
}

func (db *DB) FirstOrCreate(dest any, conds ...any) *DB {
	return &DB{db.DB.FirstOrCreate(dest, conds...)}
}

func (db *DB) FirstOrInit(dest any, conds ...any) *DB {
	return &DB{db.DB.FirstOrInit(dest, conds...)}
}

func (db *DB) Attrs(attrs ...any) *DB {
	return &DB{db.DB.Attrs(attrs...)}
}

func (db *DB) Assign(attrs ...any) *DB {
	return &DB{db.DB.Assign(attrs...)}
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

func (db *DB) Unscoped() *DB {
	return &DB{db.DB.Unscoped()}
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

func (db *DB) Transaction(fc func(tx *DB) error) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		return fc(&DB{tx})
	})
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

// 会话和调试
func (db *DB) Session(config *gorm.Session) *DB {
	return &DB{db.DB.Session(config)}
}

func (db *DB) Debug() *DB {
	return &DB{db.DB.Debug()}
}

// 错误处理
func (db *DB) Error() error {
	return errorx.TryToDBError(db.DB.Error)
}

func (db *DB) RowsAffected() int64 {
	return db.DB.RowsAffected
}

func (db *DB) Exist() bool {
	return !errors.Is(db.DB.Error, gorm.ErrRecordNotFound)
}

func (db *DB) NotExist() bool {
	return errors.Is(db.DB.Error, gorm.ErrRecordNotFound)
}

// Ping
func (db *DB) Ping() error {
	var num int
	err := db.DB.Raw("select 1").Scan(&num).Error
	return errorx.TryToDBError(err)
}

// ============ 原生SQL增删改查 ============

// QueryRow 查询单条记录
func (db *DB) QueryRow(dest any, sql string, args ...any) error {
	return errorx.TryToDBError(db.DB.Raw(sql, args...).Scan(dest).Error)
}

// QueryRows 查询多条记录
func (db *DB) QueryRows(dest any, sql string, args ...any) error {
	return errorx.TryToDBError(db.DB.Raw(sql, args...).Scan(dest).Error)
}

// InsertRow 原生插入
func (db *DB) InsertRow(sql string, args ...any) (int64, error) {
	result := db.DB.Exec(sql, args...)
	return result.RowsAffected, errorx.TryToDBError(result.Error)
}

// UpdateRaw 原生更新
func (db *DB) UpdateRaw(sql string, args ...any) (int64, error) {
	result := db.DB.Exec(sql, args...)
	return result.RowsAffected, errorx.TryToDBError(result.Error)
}

// DeleteRaw 原生删除
func (db *DB) DeleteRaw(sql string, args ...any) (int64, error) {
	result := db.DB.Exec(sql, args...)
	return result.RowsAffected, errorx.TryToDBError(result.Error)
}

// ExecRaw 执行原生SQL（通用）
func (db *DB) ExecRaw(sql string, args ...any) (int64, error) {
	result := db.DB.Exec(sql, args...)
	return result.RowsAffected, errorx.TryToDBError(result.Error)
}
