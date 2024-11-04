package model

type Ability struct {
	Id                     int    `gorm:"primaryKey"`
	Name                   string `gorm:"column:name;not null"`
	FirstAppearanceTitleId int    `gorm:"column:first_appearance_title_id;not null"`
}
