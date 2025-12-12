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

// GetByName 按名称查询
func (d *userDao) GetByName(ctx context.Context, name string) (*entity.User, error) {
	var user entity.User
	err := d.Ctx(ctx).Where(do.User{Name: name}).One(&user)
	if err != nil {
		return nil, err
	}
	if user.Id == 0 {
		return nil, nil
	}
	return &user, nil
}
