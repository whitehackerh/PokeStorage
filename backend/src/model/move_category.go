package model

type MoveCategory struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"column:name;not null"`
}