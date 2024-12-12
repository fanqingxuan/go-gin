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
	db            *gorm.DB
	ddlMigrations map[string]DDLMigration
	dmlMigrations map[string]DMLMigration
}

// NewManager 创建迁移管理器
func NewManager(db *gorm.DB) *Manager {
	return &Manager{
		db:            db,
		ddlMigrations: make(map[string]DDLMigration),
		dmlMigrations: make(map[string]DMLMigration),
	}
}

// 添加一个新的验证函数
func validateMigrationName(name string) error {
	if len(name) < 14 {
		err := fmt.Errorf("migration name '%s' is too short, must end with YYYYMMDDHHMMSS format.", name)
		color.Printf(color.Red("Migration name validation failed: %v\n"), err)
		return err
	}

	timestamp := name[len(name)-14:]
	// 验证时间戳格式
	_, err := time.Parse("20060102150405", timestamp)
	if err != nil {
		err = fmt.Errorf("migration name '%s' must end with YYYYMMDDHHMMSS format.", name)
		color.Printf(color.Red("Migration name validation failed: %v\n"), err)
		return err
	}
	return nil
}

// RegisterDDL 注册DDL迁移
func RegisterDDL(migration DDLMigration) error {
	name := reflect.TypeOf(migration).Elem().Name()
	if err := validateMigrationName(name); err != nil {
		os.Exit(0) // 使用 0 作为退出码，避免显示 exit status 1
	}
	GetManager().ddlMigrations[name] = migration
	return nil
}

// RegisterDML 注册DML迁移
func RegisterDML(migration DMLMigration) error {
	name := reflect.TypeOf(migration).Elem().Name()
	if err := validateMigrationName(name); err != nil {
		os.Exit(0) // 使用 0 作为退出码，避免显示 exit status 1
	}
	GetManager().dmlMigrations[name] = migration
	return nil
}

// initMigrationTable 初始化迁移表
func (m *Manager) initMigrationTable() error {
	if err := m.db.AutoMigrate(&Migration{}); err != nil {
		fmt.Printf(color.Red("failed to create migrations table: %v\n"), err)
		return fmt.Errorf("failed to create migrations table: %v\n", err)
	}
	return nil
}

// getExecutedMigrations 获取已执行的迁移
func (m *Manager) getExecutedMigrations() (map[string]bool, error) {
	var executedMigrations []Migration
	if err := m.db.Find(&executedMigrations).Error; err != nil {
		fmt.Printf(color.Red("failed to get executed migrations: %v\n"), err)
		return nil, fmt.Errorf("failed to get executed migrations: %v\n", err)
	}

	executedMap := make(map[string]bool)
	for _, migration := range executedMigrations {
		executedMap[migration.Desc] = true
	}
	return executedMap, nil
}

// getCurrentBatch 获取当前批次号
func (m *Manager) getCurrentBatch() (int, error) {
	var currentBatch int
	if err := m.db.Model(&Migration{}).Select("COALESCE(MAX(batch), 0)").Scan(&currentBatch).Error; err != nil {
		fmt.Printf(color.Red("failed to get current batch: %v\n"), err)
		return 0, fmt.Errorf("failed to get current batch: %v", err)
	}
	return currentBatch + 1, nil
}

// getSortedMigrationNames 获取排序后的迁移名称
func (m *Manager) getSortedMigrationNames() (ddlNames []string, dmlNames []string) {
	// 提取 DDL 迁移名称
	for name := range m.ddlMigrations {
		ddlNames = append(ddlNames, name)
	}
	// 按照时间戳后缀排序（格式：YYYYMMDDHHMMSS）
	sort.Slice(ddlNames, func(i, j int) bool {
		// 获取最后14位作为时间戳进行比较
		iTime := ddlNames[i][len(ddlNames[i])-14:]
		jTime := ddlNames[j][len(ddlNames[j])-14:]
		return iTime < jTime
	})

	// 提取 DML 迁移名称
	for name := range m.dmlMigrations {
		dmlNames = append(dmlNames, name)
	}
	// 按照时间戳后缀排序（格式：YYYYMMDDHHMMSS）
	sort.Slice(dmlNames, func(i, j int) bool {
		// 获取最后14位作为时间戳进行比较
		iTime := dmlNames[i][len(dmlNames[i])-14:]
		jTime := dmlNames[j][len(dmlNames[j])-14:]
		return iTime < jTime
	})
	return
}

// recordMigration 记录迁移
func (m *Manager) recordMigration(name string, batch int) error {
	record := Migration{
		Desc:      name,
		Batch:     batch,
		CreatedAt: time.Now(),
	}
	if err := m.db.Create(&record).Error; err != nil {
		color.Printf(color.Red("Failed to save to migration table: %s, error: %v\n"), name, err)
		return err
	}
	return nil
}

// executeDDLMigrations 执行DDL迁移
func (m *Manager) executeDDLMigrations(names []string, executedMap map[string]bool, batch int) error {
	ddlmigrator := NewDDLMigrator(m.db)
	for _, name := range names {
		if executedMap[name] {
			continue
		}

		migration := m.ddlMigrations[name]
		// 输出开始迁移
		color.Printf(color.White("Migrating: %s\n"), name)

		start := time.Now()
		if err := migration.Up(ddlmigrator); err != nil {
			color.Printf(color.Red("Failed:    %s (%v)\n"), name, err)
			return err
		}

		if err := m.recordMigration(name, batch); err != nil {
			color.Printf(color.Red("Failed to save to migration table: %s, error: %v\n"), name, err)
			return err
		}

		// 计算执行时间并输出成功信息
		duration := time.Since(start)
		color.Printf(color.Green("Migrated:  %s (%d seconds)\n"), name, int(duration.Seconds()))
	}
	return nil
}

// executeDMLMigrations 执行DML迁移
func (m *Manager) executeDMLMigrations(names []string, executedMap map[string]bool, batch int) error {
	for _, name := range names {
		if executedMap[name] {
			continue
		}

		migration := m.dmlMigrations[name]
		// 输出开始迁移
		color.Printf(color.White("Migrating: %s %s\n"), name, migration.Desc())

		start := time.Now()
		if err := migration.Handle(m.db); err != nil {
			color.Printf(color.Red("Failed:    %s (%v)\n"), name, err)
			return err
		}

		if err := m.recordMigration(name, batch); err != nil {
			color.Printf(color.Red("Failed to save to migration table: %s, error: %v\n"), name, err)
			return err
		}

		// 计算执行时间并输出成功信息
		duration := time.Since(start)
		color.Printf(color.Green("Migrated:  %s %s (%d seconds)\n"), name, migration.Desc(), int(duration.Seconds()))
	}
	return nil
}

// hasPendingMigrations 检查是否有待执行的迁移
func (m *Manager) hasPendingMigrations(ddlNames []string, dmlNames []string, executedMap map[string]bool) bool {
	// 检查DDL迁移
	for _, name := range ddlNames {
		if !executedMap[name] {
			return true
		}
	}
	// 检查DML迁移
	for _, name := range dmlNames {
		if !executedMap[name] {
			return true
		}
	}
	return false
}

// Run 执行迁移
func (m *Manager) Run() error {
	// 初始化迁移表
	if err := m.initMigrationTable(); err != nil {
		color.Printf(color.Red("Failed to initialize migration table: %v\n"), err)
		return err
	}

	// 获取已执行的迁移
	executedMap, err := m.getExecutedMigrations()
	if err != nil {
		color.Printf(color.Red("Failed to get executed migrations: %v\n"), err)
		return err
	}

	// 获取当前批次号
	currentBatch, err := m.getCurrentBatch()
	if err != nil {
		color.Printf(color.Red("Failed to get current batch: %v\n"), err)
		return err
	}

	// 获取排序后的迁移名称
	ddlNames, dmlNames := m.getSortedMigrationNames()

	// 检查是否有需要执行的迁移
	if !m.hasPendingMigrations(ddlNames, dmlNames, executedMap) {
		color.Println(color.Green("Nothing to migrate.\n"))
		return nil
	}

	// 执行DDL迁移
	if err := m.executeDDLMigrations(ddlNames, executedMap, currentBatch); err != nil {
		return err
	}

	// 执行DML迁移
	if err := m.executeDMLMigrations(dmlNames, executedMap, currentBatch); err != nil {
		return err
	}

	return nil
}
