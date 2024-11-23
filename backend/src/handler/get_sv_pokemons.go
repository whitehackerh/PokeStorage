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

func GetSVPokemons(c *gin.Context) {
	var url = "/sv/pokemons"
	var input = usecase.GetSVPokemonsInput{}

	uc := usecase.NewGetSVPokemonsInteractor(
		repository.NewSVPokemonRepository(
			infrastructure.ConnectDb(),
		),
		service.NewSVPokemonService(),
		presenter.NewGetSVPokemonsPresenter(),
	)

	output, err := uc.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
	}

	c.JSON(200, presenter.NewCommonPresenter(url, output))
}
