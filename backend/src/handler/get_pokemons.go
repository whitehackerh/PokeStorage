package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/enum"
	"github.com/whitehackerh/PokeStorage/src/infrastructure"
	"github.com/whitehackerh/PokeStorage/src/presenter"
	"github.com/whitehackerh/PokeStorage/src/repository"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

func GetPokemons(c *gin.Context) {
	var url = "/pokemons/{title-id}"
	var input = usecase.GetPokemonsInput{}

	titleId, err := strconv.Atoi(c.Param("title-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
	}

	var uc usecase.GetPokemonsUseCase
	switch titleId {
	case enum.TitleSwordShield:
		uc = usecase.NewGetSwShPokemonsInteractor(
			repository.NewSwShPokemonRepository(
				infrastructure.ConnectDb(),
			),
			service.NewSwShPokemonService(),
			presenter.NewGetPokemonsPresenter(),
		)
	case enum.TitleScarletViolet:
		uc = usecase.NewGetSVPokemonsInteractor(
			repository.NewSVPokemonRepository(
				infrastructure.ConnectDb(),
			),
			service.NewSVPokemonService(),
			presenter.NewGetPokemonsPresenter(),
		)
	default:
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
	}

	output, err := uc.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
	}

	c.JSON(200, presenter.NewCommonPresenter(url, output))
}
