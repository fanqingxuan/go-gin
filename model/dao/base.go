package dao

import (
	"context"
	"go-gin/internal/component/db"
	"go-gin/model/entity"
)

type BaseDao[T entity.Entity] struct{}

func (m *BaseDao[T]) GetById(ctx context.Context, id int64) (*T, error) {
	var e T
	pk := e.PrimaryKey()
	result := db.WithContext(ctx).Where(pk+" = ?", id).First(&e)
	if result.NotExist() {
		return nil, nil
	}
	return &e, result.Error()
}

func (m *BaseDao[T]) GetByIds(ctx context.Context, ids []int64) ([]*T, error) {
	var e T
	pk := e.PrimaryKey()
	var entities []*T
	return entities, db.WithContext(ctx).Where(pk+" IN ?", ids).Find(&entities).Error()
}

func (m *BaseDao[T]) List(ctx context.Context, query any, args ...any) ([]*T, error) {
	var entities []*T
	return entities, db.WithContext(ctx).Where(query, args...).Find(&entities).Error()
}

func (m *BaseDao[T]) Create(ctx context.Context, e *T) error {
	return db.WithContext(ctx).Create(e).Error()
}

func (m *BaseDao[T]) CreateBatch(ctx context.Context, entities []*T) error {
	return db.WithContext(ctx).Create(entities).Error()
}

func (m *BaseDao[T]) CreateBatchSize(ctx context.Context, entities []*T, batchSize int) error {
	return db.WithContext(ctx).CreateInBatches(entities, batchSize).Error()
}

func (m *BaseDao[T]) UpdateById(ctx context.Context, id int64, values any) error {
	var e T
	pk := e.PrimaryKey()
	return db.WithContext(ctx).Model(&e).Where(pk+" = ?", id).Updates(values).Error()
}

func (m *BaseDao[T]) UpdateByIds(ctx context.Context, ids []int64, values any) error {
	var e T
	pk := e.PrimaryKey()
	return db.WithContext(ctx).Model(&e).Where(pk+" IN ?", ids).Updates(values).Error()
}

func (m *BaseDao[T]) UpdateBatch(ctx context.Context, entities []*T) error {
	return db.WithContext(ctx).Save(entities).Error()
}

func (m *BaseDao[T]) DeleteById(ctx context.Context, id int64) error {
	var e T
	pk := e.PrimaryKey()
	return db.WithContext(ctx).Where(pk+" = ?", id).Delete(&e).Error()
}

func (m *BaseDao[T]) DeleteByIds(ctx context.Context, ids []int64) error {
	var e T
	pk := e.PrimaryKey()
	return db.WithContext(ctx).Where(pk+" IN ?", ids).Delete(&e).Error()
}

func (m *BaseDao[T]) Exist(ctx context.Context, id int64) (bool, error) {
	var e T
	pk := e.PrimaryKey()
	result := db.WithContext(ctx).Select(pk).Where(pk+" = ?", id).First(&e)
	if result.NotExist() {
		return false, nil
	}
	return result.Exist(), result.Error()
}

func (m *BaseDao[T]) Count(ctx context.Context, query any, args ...any) (int64, error) {
	var e T
	var count int64
	err := db.WithContext(ctx).Model(&e).Where(query, args...).Count(&count).Error()
	return count, err
}
