package model

import "time"

type JwtBlacklist struct {
	Id        string    `gorm:"primaryKey"`
	Token     string    `gorm:"column:token;not null"`
	ExpiresAt time.Time `gorm:"column:expires_at;not null"`
}
