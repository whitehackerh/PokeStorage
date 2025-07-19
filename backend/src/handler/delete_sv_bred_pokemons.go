package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitehackerh/PokeStorage/src/infrastructure"
	"github.com/whitehackerh/PokeStorage/src/presenter"
	"github.com/whitehackerh/PokeStorage/src/repository"
	"github.com/whitehackerh/PokeStorage/src/usecase"
	"gorm.io/gorm"
)

func DeleteSVBredPokemons(c *gin.Context) {
	var url = "/sv/bred-pokemons"
	var input = usecase.DeleteSVBredPokemonsInput{}

	input.BredPokemonId = c.Param("id")

	db := infrastructure.ConnectDb()
	uc := usecase.NewDeleteSVBredPokemonsInteractor(
		repository.NewSVBredPokemonRepository(db),
		repository.NewSVTeamRepository(db),
		repository.NewSVIndividualValuesRepository(db),
		repository.NewSVBasePointsRepository(db),
		repository.NewSVActualValuesRepository(db),
		presenter.NewDeleteSVBredPokemonsPresenter(),
	)

	infrastructure.WithTransaction(db, c, func(tx *gorm.DB) error {
		output, err := uc.Execute(input, tx)
		if err != nil {
			c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
			return err
		}

		c.JSON(201, presenter.NewCommonPresenter(url, output))
		return nil
	})
}
