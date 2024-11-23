package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/infrastructure"
	"github.com/whitehackerh/PokeStorage/src/presenter"
	"github.com/whitehackerh/PokeStorage/src/repository"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

func GetSwShPokemons(c *gin.Context) {
	var url = "/swsh/pokemons"
	var input = usecase.GetSwShPokemonsInput{}

	uc := usecase.NewGetSwShPokemonsInteractor(
		repository.NewSwShPokemonRepository(
			infrastructure.ConnectDb(),
		),
		service.NewSwShPokemonService(),
		presenter.NewGetSwShPokemonsPresenter(),
	)

	output, err := uc.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
	}

	c.JSON(200, presenter.NewCommonPresenter(url, output))
}
