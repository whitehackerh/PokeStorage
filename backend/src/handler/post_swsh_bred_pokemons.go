package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/infrastructure"
	"github.com/whitehackerh/PokeStorage/src/middleware"
	"github.com/whitehackerh/PokeStorage/src/presenter"
	"github.com/whitehackerh/PokeStorage/src/repository"
	"github.com/whitehackerh/PokeStorage/src/usecase"
	"gorm.io/gorm"
)

func PostSwShBredPokemons(c *gin.Context) {
	var url = "/swsh/bred-pokemons"

	claims := *c.MustGet("claims").(*middleware.Claims)

	db := infrastructure.ConnectDb()
	uc := usecase.NewPostSwShBredPokemonsInteractor(
		service.NewSwShBredPokemonService(
			service.NewBredPokemonService(),
		),
		repository.NewSwShBredPokemonRepository(db),
		repository.NewSwShIndividualValuesRepository(db),
		repository.NewSwShBasePointsRepository(db),
		repository.NewSwShActualValuesRepository(db),
		presenter.NewPostSwShBredPokemonsPresenter(),
	)

	input := usecase.PostSwShBredPokemonsInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
		return
	}

	infrastructure.WithTransaction(db, c, func(tx *gorm.DB) error {
		output, err := uc.Execute(input, claims.UserId, tx)
		if err != nil {
			c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
			return err
		}

		c.JSON(201, presenter.NewCommonPresenter(url, output))
		return nil
	})
}
