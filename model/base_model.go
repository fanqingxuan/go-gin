package model

import (
	"context"
	"go-gin/internal/component/db"
)

type Entity interface {
	PrimaryKey() string
}

// BaseEntity 提供默认主键名，可嵌入到实体结构体中
type BaseEntity struct{}

func (BaseEntity) PrimaryKey() string {
	return "id"
}

type BaseModel[T Entity] struct{}

func (m *BaseModel[T]) GetById(ctx context.Context, id int64) (*T, error) {
	var entity T
	pk := entity.PrimaryKey()
	result := db.WithContext(ctx).Where(pk+" = ?", id).First(&entity)
	if result.NotExist() {
		return nil, nil
	}
	return &entity, result.Error()
}

func (m *BaseModel[T]) GetByIds(ctx context.Context, ids []int64) ([]*T, error) {
	var entity T
	pk := entity.PrimaryKey()
	var entities []*T
	return entities, db.WithContext(ctx).Where(pk+" IN ?", ids).Find(&entities).Error()
}

func (m *BaseModel[T]) List(ctx context.Context, query any, args ...any) ([]*T, error) {
	var entities []*T
	return entities, db.WithContext(ctx).Where(query, args...).Find(&entities).Error()
}

func (m *BaseModel[T]) Create(ctx context.Context, entity *T) error {
	return db.WithContext(ctx).Create(entity).Error()
}

func (m *BaseModel[T]) CreateBatch(ctx context.Context, entities []*T) error {
	return db.WithContext(ctx).Create(entities).Error()
}

func (m *BaseModel[T]) CreateBatchSize(ctx context.Context, entities []*T, batchSize int) error {
	return db.WithContext(ctx).CreateInBatches(entities, batchSize).Error()
}

func (m *BaseModel[T]) UpdateById(ctx context.Context, id int64, values any) error {
	var entity T
	pk := entity.PrimaryKey()
	return db.WithContext(ctx).Model(&entity).Where(pk+" = ?", id).Updates(values).Error()
}

func (m *BaseModel[T]) UpdateByIds(ctx context.Context, ids []int64, values any) error {
	var entity T
	pk := entity.PrimaryKey()
	return db.WithContext(ctx).Model(&entity).Where(pk+" IN ?", ids).Updates(values).Error()
}

func (m *BaseModel[T]) UpdateBatch(ctx context.Context, entities []*T) error {
	return db.WithContext(ctx).Save(entities).Error()
}

func (m *BaseModel[T]) DeleteById(ctx context.Context, id int64) error {
	var entity T
	pk := entity.PrimaryKey()
	return db.WithContext(ctx).Where(pk+" = ?", id).Delete(&entity).Error()
}

func (m *BaseModel[T]) DeleteByIds(ctx context.Context, ids []int64) error {
	var entity T
	pk := entity.PrimaryKey()
	return db.WithContext(ctx).Where(pk+" IN ?", ids).Delete(&entity).Error()
}

func (m *BaseModel[T]) Exist(ctx context.Context, id int64) (bool, error) {
	var entity T
	pk := entity.PrimaryKey()
	result := db.WithContext(ctx).Select(pk).Where(pk+" = ?", id).First(&entity)
	if result.NotExist() {
		return false, nil
	}
	return result.Exist(), result.Error()
}

func (m *BaseModel[T]) Count(ctx context.Context, query any, args ...any) (int64, error) {
	var entity T
	var count int64
	err := db.WithContext(ctx).Model(&entity).Where(query, args...).Count(&count).Error()
	return count, err
}
