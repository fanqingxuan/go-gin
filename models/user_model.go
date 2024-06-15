package models

import (
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
