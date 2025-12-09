package model

import (
	"context"
	"go-gin/const/enum"
	"go-gin/internal/component/db"
	"time"
)

type User struct {
	BaseEntity
	Id         int64
	Name       string           `gorm:"column:name" json:"name"`
	Age        *int             `gorm:"column:age;default:null" json:"age"`
	CreateTime time.Time        `gorm:"column:create_time" json:"create_time"`
	Status     *enum.UserStatus `gorm:"column:status;default:null" json:"status"`
	UserType   *enum.UserType   `gorm:"column:user_type" json:"user_type"`
}

func (u *User) TableName() string {
	return `user`
}

type UserModel struct {
	BaseModel[User]
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

// 特有方法：指定字段创建
func (m *UserModel) Create(ctx context.Context, user *User) error {
	return db.WithContext(ctx).Select("Name", "Status", "UserType").Create(user).Error()
}

// 特有方法：按名称查询
func (m *UserModel) GetByName(ctx context.Context, name string) (*User, error) {
	var user User
	result := db.WithContext(ctx).First(&user, "name=?", name)
	if result.NotExist() {
		return nil, nil
	}
	return &user, result.Error()
}
