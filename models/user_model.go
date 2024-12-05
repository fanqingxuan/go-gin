package models

import (
	"context"
	"go-gin/internal/components/db"
	"go-gin/internal/errorx"
	"time"
)

type User struct {
	Id         int64
	Name       string    `gorm:"column:username" json:"name"`
	Age        *int      `gorm:"column:age" json:"age"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
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
	err := db.WithContext(ctx).Find(&u).Error
	return u, errorx.NewDBError(err)
}

func (m *UserModel) Add(ctx context.Context, user *User) (err error) {
	err = db.WithContext(ctx).Select("Name").Create(user).Error
	return errorx.NewDBError(err)

}

func (m *UserModel) GetByUsername(ctx context.Context, name string) (*User, error) {
	var user User
	err := db.WithContext(ctx).First(&user, "username=?", name).Error
	return &user, errorx.NewDBError(err)
}
