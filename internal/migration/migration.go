package migration

import (
	"gorm.io/gorm"
)

// Migration 迁移接口
type Migration interface {
	Up(db *gorm.DB) error
}
