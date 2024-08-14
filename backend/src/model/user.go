package model

import "time"

type User struct {
	Id           string    `gorm:"primaryKey"`
	Username     string    `gorm:"column:username"`
	Password     string    `gorm:"column:password"`
	EmailAddress string    `gorm:"column:email_address"`
	Name         string    `gorm:"column:name"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	DeletedAt    time.Time `gorm:"index"`
}
