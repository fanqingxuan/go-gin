package dao

import (
	"context"
	"go-gin/internal/component/db"
	"go-gin/model/entity"
)

type userDao struct {
	BaseDao[entity.User]
}

var User = &userDao{}

// Create 指定字段创建
func (d *userDao) Create(ctx context.Context, user *entity.User) error {
	return db.WithContext(ctx).Select("Name", "Status", "UserType").Create(user).Error()
}

// GetByName 按名称查询
func (d *userDao) GetByName(ctx context.Context, name string) (*entity.User, error) {
	var user entity.User
	result := db.WithContext(ctx).First(&user, "name=?", name)
	if result.NotExist() {
		return nil, nil
	}
	return &user, result.Error()
}
