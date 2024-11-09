package model

type Nature struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"column:name;not null"`
}
