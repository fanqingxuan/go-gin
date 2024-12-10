package migration

import (
	"gorm.io/gorm"
)

// DDLMigration DDL迁移接口
type DDLMigration interface {
	Up(db *gorm.DB) error
}

// DMLMigration DML迁移接口
type DMLMigration interface {
	Handle(db *gorm.DB) error
	Desc() string
}
