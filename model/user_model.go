package model

import (
	"context"
	"go-gin/const/enum"
	"go-gin/internal/component/db"
	"time"
)

type User struct {
	Id         int64
	Name       string           `gorm:"column:name" json:"name"`
	Age        *int             `gorm:"column:age;default:null" json:"age"`
	CreateTime time.Time        `gorm:"column:create_time" json:"create_time"`
	Status     *enum.UserStatus `gorm:"column:status;default:null" json:"status"`
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
	return u, db.WithContext(ctx).Find(&u).Error()
}

func (m *UserModel) Add(ctx context.Context, user *User) error {
	return db.WithContext(ctx).Select("Name", "Status").Create(user).Error()
}

func (m *UserModel) GetByUsername(ctx context.Context, name string) (*User, error) {
	var user User
	return &user, db.WithContext(ctx).First(&user, "username=?", name).Error()
}
