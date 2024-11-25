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

func PutSVBredPokemons(c *gin.Context) {
	var url = "/sv/bred-pokemons"

	claims := *c.MustGet("claims").(*middleware.Claims)

	db := infrastructure.ConnectDb()
	uc := usecase.NewPutSVBredPokemonsInteractor(
		service.NewSVBredPokemonService(
			service.NewBredPokemonService(),
		),
		repository.NewSVBredPokemonRepository(db),
		repository.NewSVIndividualValuesRepository(db),
		repository.NewSVBasePointsRepository(db),
		repository.NewSVActualValuesRepository(db),
		presenter.NewPutSVBredPokemonsPresenter(),
	)

	input := usecase.PutSVBredPokemonsInput{}
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
