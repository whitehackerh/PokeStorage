package repository

import (
	"gorm.io/gorm"

	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/model"
)

type TitleRepository struct {
	Db *gorm.DB
}

func NewTitleRepository(db *gorm.DB) service.ITitleRepository {
	return &TitleRepository{
		Db: db,
	}
}

func (tr *TitleRepository) Get() ([]model.Title, error) {
	var titles []model.Title
	result := tr.Db.Order("release_date DESC").Find(&titles)
	if result.Error != nil {
		return []model.Title{}, result.Error
	}
	return titles, result.Error
}
