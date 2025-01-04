package dml

import (
	"go-gin/internal/migration"

	"gorm.io/gorm"
)

func init() {
	migration.RegisterDML(&Deploy年月日时分秒{})
}

// Deploy年月日时分秒
type Deploy年月日时分秒 struct{}

// Handle 执行迁移
func (m *Deploy年月日时分秒) Handle(db *gorm.DB) error {
	return db.Exec(sql).Error
}

// Desc 获取迁移描述
func (m *Deploy年月日时分秒) Desc() string {
	return "desc"
}
