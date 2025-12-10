package entity

import (
	"go-gin/const/enum"
	"time"
)

type Entity interface {
	PrimaryKey() string
}

// BaseEntity 提供默认主键名，可嵌入到实体结构体中
type BaseEntity struct{}

func (BaseEntity) PrimaryKey() string {
	return "id"
}

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
