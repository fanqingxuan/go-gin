package model

type Xxx struct {
	Id         int64
	Name       string    `gorm:"column:snake_case" json:"snake_case"`
}

func (u *Xxx) TableName() string {
	return `tableName`
}

type XxxModel struct {
}

func NewXxxModel() *XxxModel {
	return &XxxModel{}
}