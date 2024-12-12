package dml

import (
	"go-gin/internal/migration"

	"gorm.io/gorm"
)

func init() {
	migration.RegisterDML(&Deploy20240101120000{})
}

// Deploy20240101120000001 初始化管理员用户
type Deploy20240101120000 struct{}

// Handle 执行迁移
func (m *Deploy20240101120000) Handle(db *gorm.DB) error {
	return db.Exec(`
		INSERT INTO users (username, age)
		VALUES ('admin', 18)
	`).Error
}

// Desc 获取迁移描述
func (m *Deploy20240101120000) Desc() string {
	return "用户表写入一个默认管理员用户"
}
