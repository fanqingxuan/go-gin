package enum

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type IEnum interface {
	Code() int
	Desc() string
	String() string
	// 通过 sql.Scanner 和 sql.Valuer 接口实现数据库的存储和读取
	// 这两个接口是 go-sql-driver/mysql 包提供的
	sql.Scanner
	driver.Valuer
	json.Marshaler
	json.Unmarshaler
}
