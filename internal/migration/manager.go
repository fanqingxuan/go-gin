package migration

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"time"

	"github.com/labstack/gommon/color"
	"gorm.io/gorm"
)

// Manager 迁移管理器
type Manager struct {
	db         *gorm.DB
	migrations map[string]Migration
}

// NewManager 创建迁移管理器
func NewManager(db *gorm.DB) *Manager {
	return &Manager{
		db:         db,
		migrations: make(map[string]Migration),
	}
}

// validateMigrationName 验证迁移名称格式
func validateMigrationName(name string) error {
	if len(name) < 14 {
		err := fmt.Errorf("migration name '%s' is too short, must end with YYYYMMDDHHMMSS format", name)
		color.Printf(color.Red("Migration name validation failed: %v\n"), err)
		return err
	}

	timestamp := name[len(name)-14:]
	if _, err := time.Parse("20060102150405", timestamp); err != nil {
		err = fmt.Errorf("migration name '%s' must end with YYYYMMDDHHMMSS format", name)
		color.Printf(color.Red("Migration name validation failed: %v\n"), err)
		return err
	}
	return nil
}

// Register 注册迁移
func Register(m Migration) {
	name := reflect.TypeOf(m).Elem().Name()
	if err := validateMigrationName(name); err != nil {
		os.Exit(1)
	}
	GetManager().migrations[name] = m
}

// initMigrationTable 初始化迁移表
func (m *Manager) initMigrationTable() error {
	return m.db.AutoMigrate(&MigrationRecord{})
}

// getExecutedMigrations 获取已执行的迁移
func (m *Manager) getExecutedMigrations() (map[string]bool, error) {
	var records []MigrationRecord
	if err := m.db.Find(&records).Error; err != nil {
		return nil, err
	}

	executed := make(map[string]bool)
	for _, r := range records {
		executed[r.Migration] = true
	}
	return executed, nil
}

// getCurrentBatch 获取当前批次号
func (m *Manager) getCurrentBatch() (int, error) {
	var batch int
	err := m.db.Model(&MigrationRecord{}).Select("COALESCE(MAX(batch), 0)").Scan(&batch).Error
	return batch + 1, err
}

// getSortedNames 获取排序后的迁移名称
func (m *Manager) getSortedNames() []string {
	names := make([]string, 0, len(m.migrations))
	for name := range m.migrations {
		names = append(names, name)
	}
	sort.Slice(names, func(i, j int) bool {
		return names[i][len(names[i])-14:] < names[j][len(names[j])-14:]
	})
	return names
}

// Run 执行迁移
func (m *Manager) Run() error {
	if err := m.initMigrationTable(); err != nil {
		color.Printf(color.Red("Failed to initialize migration table: %v\n"), err)
		return err
	}

	executed, err := m.getExecutedMigrations()
	if err != nil {
		color.Printf(color.Red("Failed to get executed migrations: %v\n"), err)
		return err
	}

	batch, err := m.getCurrentBatch()
	if err != nil {
		color.Printf(color.Red("Failed to get current batch: %v\n"), err)
		return err
	}

	names := m.getSortedNames()
	pending := 0
	for _, name := range names {
		if !executed[name] {
			pending++
		}
	}

	if pending == 0 {
		color.Println(color.Green("Nothing to migrate."))
		return nil
	}

	for _, name := range names {
		if executed[name] {
			continue
		}

		color.Printf(color.White("Migrating: %s\n"), name)
		start := time.Now()

		if err := m.migrations[name].Up(m.db); err != nil {
			color.Printf(color.Red("Failed:    %s (%v)\n"), name, err)
			return err
		}

		record := MigrationRecord{
			Migration: name,
			Batch:     batch,
		}
		if err := m.db.Create(&record).Error; err != nil {
			color.Printf(color.Red("Failed to record migration: %s (%v)\n"), name, err)
			return err
		}

		duration := time.Since(start)
		color.Printf(color.Green("Migrated:  %s (%.2fs)\n"), name, duration.Seconds())
	}

	return nil
}
