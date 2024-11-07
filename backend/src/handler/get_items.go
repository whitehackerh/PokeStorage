package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/whitehackerh/PokeStorage/src/enum"
	"github.com/whitehackerh/PokeStorage/src/infrastructure"
	"github.com/whitehackerh/PokeStorage/src/presenter"
	"github.com/whitehackerh/PokeStorage/src/repository"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

func GetItems(c *gin.Context) {
	var url = "/items"
	var input = usecase.GetItemsInput{}

	titleId, err := strconv.Atoi(c.Param("title-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
	}

	var uc usecase.GetItemsUseCase
	switch titleId {
	case enum.TitleSwordShield:
		uc = usecase.NewSwShGetItemsInteractor(
			repository.NewSwShItemRepository(
				infrastructure.ConnectDb(),
			),
			presenter.NewGetItemsPresenter(),
		)
	case enum.TitleScarletViolet:
		uc = usecase.NewSVGetItemsInteractor(
			repository.NewSVItemRepository(
				infrastructure.ConnectDb(),
			),
			presenter.NewGetItemsPresenter(),
		)
	default:
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
	}

	output, err := uc.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
		return
	}

	c.JSON(200, presenter.NewCommonPresenter(url, output))
}
