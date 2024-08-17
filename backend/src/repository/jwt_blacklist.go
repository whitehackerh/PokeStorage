package repository

import (
	"github.com/whitehackerh/PokeStorage/src/middleware"
	"github.com/whitehackerh/PokeStorage/src/model"
	"gorm.io/gorm"
)

type JwtBlacklistRepository struct {
	Db *gorm.DB
}

func NewJwtBlacklistRepository(db *gorm.DB) middleware.IJwtBlacklistRepository {
	return &JwtBlacklistRepository{
		Db: db,
	}
}

func (jbr *JwtBlacklistRepository) FindByToken(token string) (model.JwtBlacklist, error) {
	var jwtBlacklist model.JwtBlacklist
	result := jbr.Db.Where("token = ?", token).First(&jwtBlacklist)
	if result.Error != nil {
		return model.JwtBlacklist{}, result.Error
	}
	return jwtBlacklist, result.Error
}

func (jbr *JwtBlacklistRepository) Create(jwtBlacklist model.JwtBlacklist) error {
	result := jbr.Db.Create(&jwtBlacklist)
	return result.Error
}
