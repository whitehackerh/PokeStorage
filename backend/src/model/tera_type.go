package model

type TeraType struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"column:name;not null"`
}
