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
		CREATE TABLE users12 (
			id bigint unsigned NOT NULL AUTO_INCREMENT,
			name varchar(255) NOT NULL,
			email varchar(255) NOT NULL,
			created_at timestamp NULL DEFAULT NULL,
			updated_at timestamp NULL DEFAULT NULL,
			PRIMARY KEY (id),
			UNIQUE KEY users_email_unique (email)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	`).Error
}
