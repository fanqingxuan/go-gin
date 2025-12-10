package entityx

type Entity interface {
	PrimaryKey() string
}

// BaseEntity 提供默认主键名，可嵌入到实体结构体中
type BaseEntity struct{}

func (BaseEntity) PrimaryKey() string {
	return "id"
}
