package model

type SVItem struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"column:name;not null"`
}
