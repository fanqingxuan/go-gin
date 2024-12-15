package model

import (
	"context"
	"go-gin/internal/components/db"
	"time"
)

type User struct {
	Id         int64
	Name       string    `gorm:"column:username" json:"name"`
	Age        *int      `gorm:"column:age" json:"age"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}

func (u *User) TableName() string {
	return `users`
}

type UserModel struct {
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (m *UserModel) List(ctx context.Context) ([]User, error) {
	var u []User
	return u, db.WithContext(ctx).Find(&u).Error()
}

func (m *UserModel) Add(ctx context.Context, user *User) error {
	return db.WithContext(ctx).Select("Name").Create(user).Error()
}

func (m *UserModel) GetByUsername(ctx context.Context, name string) (*User, error) {
	var user User
	return &user, db.WithContext(ctx).First(&user, "username=?", name).Error()
}
