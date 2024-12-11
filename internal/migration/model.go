package migration

import (
	"time"
)

// Migration 迁移记录模型
type Migration struct {
	ID        uint      `gorm:"primaryKey"`
	Desc      string    `gorm:"column:desc;type:varchar(255);not null"`
	Batch     int       `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null"`
}

// TableName 指定表名
func (Migration) TableName() string {
	return "migrations"
}
