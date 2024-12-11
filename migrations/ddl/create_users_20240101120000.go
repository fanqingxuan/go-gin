package ddl

import (
	"go-gin/internal/migration"

	"gorm.io/gorm"
)

func init() {
	migration.RegisterDDL(&CreateUsers20240101120000{})
}

// CreateUsers20240101120000 创建用户表迁移
type CreateUsers20240101120000 struct{}

// Up 执行迁移
func (m *CreateUsers20240101120000) Up(db *gorm.DB) error {
	return db.Exec(`
		CREATE TABLE users (
			id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(255) DEFAULT '' COMMENT '用户名',
			age INT DEFAULT 0 COMMENT '年龄',
			create_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'
		)
	`).Error
}
