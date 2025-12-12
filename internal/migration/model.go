package migration

import "time"

// MigrationRecord 迁移记录模型
type MigrationRecord struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Migration string    `gorm:"column:migration;type:varchar(255);not null;uniqueIndex"`
	Batch     int       `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

// TableName 指定表名
func (MigrationRecord) TableName() string {
	return "migrations"
}
