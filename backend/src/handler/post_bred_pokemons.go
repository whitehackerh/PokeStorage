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

func PostBredPokemons(c *gin.Context) {
	var url = "/bred-pokemons/{title-id}"

	titleId, err := strconv.Atoi(c.Param("title-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
	}
	claims := *c.MustGet("claims").(*middleware.Claims)

	db := infrastructure.ConnectDb()
	switch titleId {
	case enum.TitleSwordShield:
		PostSwShBredPokemons(db, c, url, claims.UserId)
	case enum.TitleScarletViolet:
		PostSVBredPokemons(db, c, url, claims.UserId)
	}
}

func PostSwShBredPokemons(db *gorm.DB, c *gin.Context, url string, userId string) {
	uc := usecase.NewPostSwShBredPokemonsInteractor(
		service.NewSwShBredPokemonService(
			service.NewBredPokemonService(),
		),
		repository.NewSwShBredPokemonRepository(db),
		repository.NewSwShIndividualValuesRepository(db),
		repository.NewSwShBasePointsRepository(db),
		repository.NewSwShActualValuesRepository(db),
		presenter.NewPostBredPokemonsPresenter(),
	)

	input := usecase.PostSwShBredPokemonsInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
		return
	}

	infrastructure.WithTransaction(db, c, func(tx *gorm.DB) error {
		output, err := uc.Execute(input, userId, tx)
		if err != nil {
			c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
			return err
		}

		c.JSON(201, presenter.NewCommonPresenter(url, output))
		return nil
	})
}

func PostSVBredPokemons(db *gorm.DB, c *gin.Context, url string, userId string) {
	uc := usecase.NewPostSVBredPokemonsInteractor(
		service.NewSVBredPokemonService(
			service.NewBredPokemonService(),
		),
		repository.NewSVBredPokemonRepository(db),
		repository.NewSVIndividualValuesRepository(db),
		repository.NewSVBasePointsRepository(db),
		repository.NewSVActualValuesRepository(db),
		presenter.NewPostBredPokemonsPresenter(),
	)

	input := usecase.PostSVBredPokemonsInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
		return
	}

	infrastructure.WithTransaction(db, c, func(tx *gorm.DB) error {
		output, err := uc.Execute(input, userId, tx)
		if err != nil {
			c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
			return err
		}

		c.JSON(201, presenter.NewCommonPresenter(url, output))
		return nil
	})
}
