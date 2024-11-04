package model

type SwShItem struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"column:name;not null"`
}
