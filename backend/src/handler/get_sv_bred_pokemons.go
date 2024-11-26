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
)

func GetSVBredPokemons(c *gin.Context) {
	var url = "/sv/bred-pokemons"

	claims := *c.MustGet("claims").(*middleware.Claims)

	uc := usecase.NewGetSVBredPokemonsInteractor(
		service.NewSVBredPokemonService(
			service.NewBredPokemonService(),
		),
		repository.NewSVBredPokemonRepository(infrastructure.ConnectDb()),
		presenter.NewGetSVBredPokemonsPresenter(),
	)

	output, err := uc.Execute(usecase.GetSVBredPokemonsInput{
		UserId: claims.UserId,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
		return
	}

	c.JSON(200, presenter.NewCommonPresenter(url, output))
}
