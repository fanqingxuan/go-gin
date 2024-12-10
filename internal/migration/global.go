package migration

import (
	"sync"

	"gorm.io/gorm"
)

var (
	manager *Manager
	once    sync.Once
)

// GetManager 获取全局迁移管理器
func GetManager() *Manager {
	once.Do(func() {
		manager = &Manager{
			ddlMigrations: make(map[string]DDLMigration),
			dmlMigrations: make(map[string]DMLMigration),
		}
	})
	return manager
}

// SetDB 设置数据库连接
func SetDB(db *gorm.DB) {
	GetManager().db = db
}
