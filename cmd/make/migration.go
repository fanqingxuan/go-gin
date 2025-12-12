package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"
)

const migrationTmpl = `package migration

import (
	"go-gin/internal/migration"

	"gorm.io/gorm"
)

func init() {
	migration.Register(&{{.StructName}}{})
}

// {{.StructName}} {{.Description}}
type {{.StructName}} struct{}

func (m *{{.StructName}}) Up(db *gorm.DB) error {
	return db.Exec(` + "`" + `
		-- TODO: 编写迁移 SQL
	` + "`" + `).Error
}
`

type MigrationData struct {
	StructName  string
	Description string
}

func runMigration(args []string) {
	if len(args) < 1 {
		fmt.Println("用法: go run ./cmd/make/... make:migration <name>")
		fmt.Println("示例: go run ./cmd/make/... make:migration create_orders")
		os.Exit(1)
	}

	name := args[0]
	timestamp := time.Now().Format("20060102150405")
	fileName := fmt.Sprintf("migration/%s_%s.go", strings.ToLower(strings.ReplaceAll(name, "-", "_")), timestamp)
	structName := toPascalCase(name) + timestamp

	data := MigrationData{
		StructName:  structName,
		Description: name,
	}

	tmpl := template.Must(template.New("migration").Parse(migrationTmpl))

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("创建文件失败: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	if err := tmpl.Execute(file, data); err != nil {
		fmt.Printf("生成文件失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("已创建: %s\n", fileName)
}
