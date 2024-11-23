package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitehackerh/PokeStorage/src/infrastructure"
	"github.com/whitehackerh/PokeStorage/src/presenter"
	"github.com/whitehackerh/PokeStorage/src/repository"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

func GetSVItems(c *gin.Context) {
	var url = "/sv/items"
	var input = usecase.GetSVItemsInput{}

	uc := usecase.NewSVGetItemsInteractor(
		repository.NewSVItemRepository(
			infrastructure.ConnectDb(),
		),
		presenter.NewGetSVItemsPresenter(),
	)

	output, err := uc.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
		return
	}

	c.JSON(200, presenter.NewCommonPresenter(url, output))
}
