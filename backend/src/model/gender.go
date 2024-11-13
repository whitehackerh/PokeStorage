package model

type Gender struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"column:name;not null"`
}
