// 生成迁移文件
// 用法: go run cmd/make_migration/main.go create_orders
//       go run cmd/make_migration/main.go add_email_to_users
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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: go run cmd/make_migration/main.go <migration_name>")
		fmt.Println("示例: go run cmd/make_migration/main.go create_orders")
		os.Exit(1)
	}

	name := os.Args[1]
	timestamp := time.Now().Format("20060102150405")
	fileName := fmt.Sprintf("migration/%s_%s.go", toSnakeCase(name), timestamp)
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

func toPascalCase(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(part[:1]) + part[1:]
		}
	}
	return strings.Join(parts, "")
}

func toSnakeCase(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, "-", "_"))
}
