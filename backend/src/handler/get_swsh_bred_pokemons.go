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

func GetSwShBredPokemons(c *gin.Context) {
	var url = "/swsh/bred-pokemons"

	claims := *c.MustGet("claims").(*middleware.Claims)

	uc := usecase.NewGetSwShBredPokemonsInteractor(
		service.NewSwShBredPokemonService(
			service.NewBredPokemonService(),
		),
		repository.NewSwShBredPokemonRepository(infrastructure.ConnectDb()),
		presenter.NewGetSwShBredPokemonsPresenter(),
	)

	output, err := uc.Execute(usecase.GetSwShBredPokemonsInput{
		UserId: claims.UserId,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
		return
	}

	c.JSON(200, presenter.NewCommonPresenter(url, output))
}
