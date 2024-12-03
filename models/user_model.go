package models

import (
	"context"
	"go-gin/internal/components/db"
	"time"
)

type User struct {
	Id         int64
	Name       string    `gorm:"column:username"`
	Age        *int      `gorm:"column:age"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func (u *User) TableName() string {
	return `user`
}

type UserModel struct {
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (m *UserModel) List(ctx context.Context) ([]User, error) {
	var u []User
	if err := db.WithContext(ctx).Find(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (m *UserModel) Add(ctx context.Context, user *User) (err error) {
	if err = db.WithContext(ctx).Select("Name").Create(user).Error; err != nil {
		return
	}
	return nil
}

func (m *UserModel) GetByUsername(ctx context.Context, name string) (*User, error) {
	var user User
	if err := db.WithContext(ctx).First(&user, "username=?", name).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
