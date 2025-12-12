package db

import (
	"context"
	"database/sql"

	"go-gin/internal/errorx"

	"gorm.io/gorm"
)

// ============ 数据库连接 ============

// Ping 检查数据库连接
func Ping(ctx context.Context) error {
	sqlDB, err := instance.DB()
	if err != nil {
		return errorx.TryToDBError(err)
	}
	return errorx.TryToDBError(sqlDB.PingContext(ctx))
}

// WithContext 返回带上下文的数据库操作对象（兼容旧 API）
func WithContext(ctx context.Context) *ContextDB {
	return &ContextDB{ctx: ctx, db: instance.WithContext(ctx)}
}

// ContextDB 带上下文的数据库操作
type ContextDB struct {
	ctx context.Context
	db  *gorm.DB
}

// Raw 执行原生 SQL
func (c *ContextDB) Raw(sql string, args ...any) *RawModel {
	return &RawModel{ctx: c.ctx, db: c.db, sql: sql, args: args}
}

// Ping 检查数据库连接
func (c *ContextDB) Ping() error {
	sqlDB, err := c.db.DB()
	if err != nil {
		return errorx.TryToDBError(err)
	}
	return errorx.TryToDBError(sqlDB.PingContext(c.ctx))
}

// ============ 原生 SQL 操作 ============

// Raw 基于原生 SQL 创建 Model
func Raw(ctx context.Context, sql string, args ...any) *RawModel {
	return &RawModel{
		ctx:  ctx,
		db:   instance.WithContext(ctx),
		sql:  sql,
		args: args,
	}
}

// ============ 原生 SQL 查询 ============

// Query 执行查询 SQL，返回多行结果
func Query(ctx context.Context, sql string, args ...any) ([]map[string]any, error) {
	var result []map[string]any
	err := instance.WithContext(ctx).Raw(sql, args...).Scan(&result).Error
	return result, errorx.TryToDBError(err)
}

// Exec 执行非查询 SQL（INSERT/UPDATE/DELETE）
func Exec(ctx context.Context, sql string, args ...any) (sql.Result, error) {
	result := instance.WithContext(ctx).Exec(sql, args...)
	return &execResult{
		rowsAffected: result.RowsAffected,
		err:          result.Error,
	}, errorx.TryToDBError(result.Error)
}

// ============ 便捷查询方法 ============

// GetAll 查询多条记录
func GetAll(ctx context.Context, sql string, args ...any) ([]map[string]any, error) {
	return Query(ctx, sql, args...)
}

// GetOne 查询单条记录
func GetOne(ctx context.Context, sql string, args ...any) (map[string]any, error) {
	var result map[string]any
	err := instance.WithContext(ctx).Raw(sql, args...).Scan(&result).Error
	return result, errorx.TryToDBError(err)
}

// GetValue 查询单个值
func GetValue(ctx context.Context, sql string, args ...any) (any, error) {
	var result map[string]any
	err := instance.WithContext(ctx).Raw(sql, args...).Scan(&result).Error
	if err != nil {
		return nil, errorx.TryToDBError(err)
	}
	for _, v := range result {
		return v, nil
	}
	return nil, nil
}

// GetCount 查询数量
func GetCount(ctx context.Context, sql string, args ...any) (int64, error) {
	var count int64
	err := instance.WithContext(ctx).Raw(sql, args...).Scan(&count).Error
	return count, errorx.TryToDBError(err)
}

// GetArray 查询单列返回数组
func GetArray(ctx context.Context, sql string, args ...any) ([]any, error) {
	var results []map[string]any
	err := instance.WithContext(ctx).Raw(sql, args...).Scan(&results).Error
	if err != nil {
		return nil, errorx.TryToDBError(err)
	}
	arr := make([]any, 0, len(results))
	for _, row := range results {
		for _, v := range row {
			arr = append(arr, v)
			break
		}
	}
	return arr, nil
}

// GetScan 查询并扫描到指定结构
func GetScan(ctx context.Context, dest any, sql string, args ...any) error {
	return errorx.TryToDBError(instance.WithContext(ctx).Raw(sql, args...).Scan(dest).Error)
}

// ============ CRUD 快捷方法 ============

// Insert 插入数据
func Insert(ctx context.Context, table string, data any) (sql.Result, error) {
	result := instance.WithContext(ctx).Table(table).Create(data)
	return &execResult{rowsAffected: result.RowsAffected, err: result.Error}, errorx.TryToDBError(result.Error)
}

// InsertAndGetId 插入并返回自增 ID
func InsertAndGetId(ctx context.Context, table string, data any) (int64, error) {
	result := instance.WithContext(ctx).Table(table).Create(data)
	if result.Error != nil {
		return 0, errorx.TryToDBError(result.Error)
	}
	return result.RowsAffected, nil
}

// Update 更新数据
func Update(ctx context.Context, table string, data any, condition string, args ...any) (sql.Result, error) {
	result := instance.WithContext(ctx).Table(table).Where(condition, args...).Updates(data)
	return &execResult{rowsAffected: result.RowsAffected, err: result.Error}, errorx.TryToDBError(result.Error)
}

// Delete 删除数据
func Delete(ctx context.Context, table string, condition string, args ...any) (sql.Result, error) {
	result := instance.WithContext(ctx).Table(table).Where(condition, args...).Delete(nil)
	return &execResult{rowsAffected: result.RowsAffected, err: result.Error}, errorx.TryToDBError(result.Error)
}

// ============ 事务 ============

// Begin 开始事务
func Begin(ctx context.Context) *TX {
	tx := instance.WithContext(ctx).Begin()
	return &TX{db: tx, ctx: ctx}
}

// Transaction 事务执行（自动提交/回滚）
func Transaction(ctx context.Context, fc func(tx *TX) error) error {
	return errorx.TryToDBError(instance.WithContext(ctx).Transaction(func(gormTx *gorm.DB) error {
		return fc(&TX{db: gormTx, ctx: ctx})
	}))
}

// ============ 事务结构 ============

// TX 事务结构
type TX struct {
	db  *gorm.DB
	ctx context.Context
}

// Model 在事务中创建 Model
func (tx *TX) Model(table string) *Model {
	return &Model{
		ctx:        tx.ctx,
		db:         tx.db,
		table:      table,
		primaryKey: "id",
	}
}

// Raw 在事务中执行原生 SQL
func (tx *TX) Raw(sql string, args ...any) *RawModel {
	return &RawModel{
		ctx:  tx.ctx,
		db:   tx.db,
		sql:  sql,
		args: args,
	}
}

// Exec 在事务中执行 SQL
func (tx *TX) Exec(sql string, args ...any) (sql.Result, error) {
	result := tx.db.Exec(sql, args...)
	return &execResult{rowsAffected: result.RowsAffected, err: result.Error}, errorx.TryToDBError(result.Error)
}

// Commit 提交事务
func (tx *TX) Commit() error {
	return errorx.TryToDBError(tx.db.Commit().Error)
}

// Rollback 回滚事务
func (tx *TX) Rollback() error {
	return errorx.TryToDBError(tx.db.Rollback().Error)
}

// ============ RawModel 原生 SQL 模型 ============

// RawModel 原生 SQL 模型
type RawModel struct {
	ctx  context.Context
	db   *gorm.DB
	sql  string
	args []any
}

// Scan 扫描结果到目标
func (r *RawModel) Scan(dest any) *RawModel {
	r.db = r.db.Raw(r.sql, r.args...).Scan(dest)
	return r
}

// Error 返回错误
func (r *RawModel) Error() error {
	return errorx.TryToDBError(r.db.Error)
}

// Exec 执行 SQL
func (r *RawModel) Exec() (sql.Result, error) {
	result := r.db.Exec(r.sql, r.args...)
	return &execResult{rowsAffected: result.RowsAffected, err: result.Error}, errorx.TryToDBError(result.Error)
}

// ============ 辅助结构 ============

type execResult struct {
	rowsAffected int64
	err          error
}

func (r *execResult) LastInsertId() (int64, error) {
	return 0, nil // GORM 不直接支持
}

func (r *execResult) RowsAffected() (int64, error) {
	return r.rowsAffected, r.err
}
