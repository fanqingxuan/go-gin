package migration

import (
	"gorm.io/gorm"
)

// DDLMigrator DDL迁移器
type DDLMigrator struct {
	db *gorm.DB
}

// NewDDLMigrator 创建DDL迁移器
func NewDDLMigrator(db *gorm.DB) *DDLMigrator {
	return &DDLMigrator{db: db}
}

// DropTable 删除表
func (m *DDLMigrator) DropTable(tablename string) error {
	return m.db.Migrator().DropTable(tablename)
}

// HasTable 检查表是否存在
func (m *DDLMigrator) HasTable(tablename string) bool {
	return m.db.Migrator().HasTable(tablename)
}

// RenameTable 重命名表
func (m *DDLMigrator) RenameTable(oldtablename, newtablename string) error {
	return m.db.Migrator().RenameTable(oldtablename, newtablename)
}

// DropColumn 删除列
func (m *DDLMigrator) DropColumn(tablename, columnname string) error {
	return m.db.Migrator().DropColumn(tablename, columnname)
}

// HasColumn 检查列是否存在
func (m *DDLMigrator) HasColumn(tablename, columnname string) bool {
	return m.db.Migrator().HasColumn(tablename, columnname)
}

// RenameColumn 重命名列
func (m *DDLMigrator) RenameColumn(tablename, oldcolumnname, newcolumnname string) error {
	return m.db.Migrator().RenameColumn(tablename, oldcolumnname, newcolumnname)
}

// CreateIndex 创建索引
func (m *DDLMigrator) CreateIndex(tablename, indexname string) error {
	return m.db.Migrator().CreateIndex(tablename, indexname)
}

// DropIndex 删除索引
func (m *DDLMigrator) DropIndex(tablename, indexname string) error {
	return m.db.Migrator().DropIndex(tablename, indexname)
}

// HasIndex 检查索引是否存在
func (m *DDLMigrator) HasIndex(tablename, indexname string) bool {
	return m.db.Migrator().HasIndex(tablename, indexname)
}

// RenameIndex 重命名索引
func (m *DDLMigrator) RenameIndex(tablename, oldindexname, newindexname string) error {
	return m.db.Migrator().RenameIndex(tablename, oldindexname, newindexname)
}

// DMLMigrator DML迁移器
type DMLMigrator struct {
	db *gorm.DB
}

// NewDMLMigrator 创建DML迁移器
func NewDMLMigrator(db *gorm.DB) *DMLMigrator {
	return &DMLMigrator{db: db}
}

// DB 获取数据库连接
func (m *DMLMigrator) DB() *gorm.DB {
	return m.db
}
