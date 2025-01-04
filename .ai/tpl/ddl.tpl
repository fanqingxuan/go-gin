package ddl

import (
	"go-gin/internal/migration"
)

func init() {
	migration.RegisterDDL(&CreateXxx年月日时分秒{})
}

// CreateXxx年月日时分秒
type CreateXxx年月日时分秒 struct{}

// Up 执行迁移
func (m *CreateXxx年月日时分秒) Up(migrator *migration.DDLMigrator) error {
    return migrator.Exec(sql)
}
