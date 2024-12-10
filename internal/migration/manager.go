package migration

import (
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"time"

	"github.com/labstack/gommon/color"
	"gorm.io/gorm"
)

// 文件命名规则
var (
	ddlFilePattern = regexp.MustCompile(`^[a-zA-Z]+\d{14}$`)
	dmlFilePattern = regexp.MustCompile(`^Deploy\d{14}$`)
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

// RegisterDDL 注册DDL迁移
func RegisterDDL(migration DDLMigration) error {
	name := reflect.TypeOf(migration).Elem().Name()
	if !ddlFilePattern.MatchString(name) {
		fmt.Printf(color.Red("invalid DDL migration name format: %s\n"), name)
		return fmt.Errorf("invalid DDL migration name format: %s", name)
	}
	GetManager().ddlMigrations[name] = migration
	return nil
}

// RegisterDML 注册DML迁移
func RegisterDML(migration DMLMigration) error {
	name := reflect.TypeOf(migration).Elem().Name()
	if !dmlFilePattern.MatchString(name) {
		fmt.Printf(color.Red("invalid DML migration name format: %s\n"), name)
		return fmt.Errorf("invalid DML migration name format: %s", name)
	}
	GetManager().dmlMigrations[name] = migration
	return nil
}

// initMigrationTable 初始化迁移表
func (m *Manager) initMigrationTable() error {
	if err := m.db.AutoMigrate(&Migration{}); err != nil {
		return fmt.Errorf("failed to create migrations table: %v", err)
	}
	return nil
}

// getExecutedMigrations 获取已执行的迁移
func (m *Manager) getExecutedMigrations() (map[string]bool, error) {
	var executedMigrations []Migration
	if err := m.db.Find(&executedMigrations).Error; err != nil {
		return nil, fmt.Errorf("failed to get executed migrations: %v", err)
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
		return 0, fmt.Errorf("failed to get current batch: %v", err)
	}
	return currentBatch + 1, nil
}

// getSortedMigrationNames 获取排序后的迁移名称
func (m *Manager) getSortedMigrationNames() (ddlNames []string, dmlNames []string) {
	for name := range m.ddlMigrations {
		ddlNames = append(ddlNames, name)
	}
	sort.Strings(ddlNames)

	for name := range m.dmlMigrations {
		dmlNames = append(dmlNames, name)
	}
	sort.Strings(dmlNames)
	return
}

// recordMigration 记录迁移
func (m *Manager) recordMigration(name string, batch int) error {
	record := Migration{
		Desc:  name,
		Batch: batch,
	}
	if err := m.db.Create(&record).Error; err != nil {
		color.Printf(color.Red("Failed to save to migration table: %s, error: %v\n"), name, err)
		return err
	}
	return nil
}

// executeDDLMigrations 执行DDL迁移
func (m *Manager) executeDDLMigrations(names []string, executedMap map[string]bool, batch int) error {
	color.Println(color.White("Start executing DDL migrations..."))
	for _, name := range names {
		if executedMap[name] {
			continue
		}

		color.Printf(color.White("Executing DDL migration: %s\n"), name)

		if err := m.ddlMigrations[name].Up(m.db); err != nil {
			color.Printf(color.Red("Migration failed: %s, error: %v\n"), name, err)
			return err
		}

		if err := m.recordMigration(name, batch); err != nil {
			return err
		}
		color.Printf(color.Green("Migration successful: %s\n"), name)
	}
	return nil
}

// executeDMLMigrations 执行DML迁移
func (m *Manager) executeDMLMigrations(names []string, executedMap map[string]bool, batch int) error {
	color.Println(color.White("Start executing DML migrations..."))
	for _, name := range names {
		if executedMap[name] {
			continue
		}

		color.Printf(color.White("Executing DML migration: %s\n"), name)

		if err := m.dmlMigrations[name].Handle(m.db); err != nil {
			color.Printf(color.Red("Migration failed: %s, error: %v\n"), name, err)
			return err
		}

		if err := m.recordMigration(name, batch); err != nil {
			return err
		}
		color.Printf(color.Green("Migration successful: %s\n"), name)
	}
	return nil
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

// GenerateMigrationName 生成迁移文件名
func GenerateMigrationName(prefix string) string {
	timestamp := time.Now().Format("20060102150405")
	// Generate sequence number 001, you may need to adjust this based on existing files
	return fmt.Sprintf("%s_%s%s", prefix, timestamp, "001")
}
