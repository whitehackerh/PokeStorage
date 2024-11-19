package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/enum"
	"github.com/whitehackerh/PokeStorage/src/infrastructure"
	"github.com/whitehackerh/PokeStorage/src/middleware"
	"github.com/whitehackerh/PokeStorage/src/presenter"
	"github.com/whitehackerh/PokeStorage/src/repository"
	"github.com/whitehackerh/PokeStorage/src/usecase"
	"gorm.io/gorm"
)

func GetBredPokemons(c *gin.Context) {
	var url = "/bred-pokemons/{title-id}"

	titleId, err := strconv.Atoi(c.Param("title-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
	}
	claims := *c.MustGet("claims").(*middleware.Claims)

	db := infrastructure.ConnectDb()
	switch titleId {
	case enum.TitleSwordShield:
		GetSwShBredPokemons(db, c, url, claims.UserId)
	case enum.TitleScarletViolet:
		GetSVBredPokemons(db, c, url, claims.UserId)
	}
}

func GetSwShBredPokemons(db *gorm.DB, c *gin.Context, url string, userId string) {
	uc := usecase.NewGetSwShBredPokemonsInteractor(
		service.NewSwShBredPokemonService(
			service.NewBredPokemonService(),
		),
		repository.NewSwShBredPokemonRepository(db),
		presenter.NewGetSwShBredPokemonsPresenter(),
	)

	output, err := uc.Execute(usecase.GetSwShBredPokemonsInput{
		UserId: userId,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
		return
	}

	c.JSON(200, presenter.NewCommonPresenter(url, output))
}

func GetSVBredPokemons(db *gorm.DB, c *gin.Context, url string, userId string) {
	uc := usecase.NewGetSVBredPokemonsInteractor(
		service.NewSVBredPokemonService(
			service.NewBredPokemonService(),
		),
		repository.NewSVBredPokemonRepository(db),
		presenter.NewGetSVBredPokemonsPresenter(),
	)

	output, err := uc.Execute(usecase.GetSVBredPokemonsInput{
		UserId: userId,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
		return
	}

	c.JSON(200, presenter.NewCommonPresenter(url, output))
}
