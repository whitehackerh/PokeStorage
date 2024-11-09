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

func GetMoves(c *gin.Context) {
	var url = "moves"
	var input = usecase.GetMovesInput{}

	titleId, err := strconv.Atoi(c.Param("title-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
	}

	var uc usecase.GetMovesUseCase
	switch titleId {
	case enum.TitleSwordShield:
		uc = usecase.NewSwShGetMovesInteractor(
			repository.NewSwShMoveRepository(
				infrastructure.ConnectDb(),
			),
			service.NewSwShMoveService(),
			presenter.NewGetMovesPresenter(),
		)
	case enum.TitleScarletViolet:
		uc = usecase.NewSVGetMovesInteractor(
			repository.NewSVMoveRepository(
				infrastructure.ConnectDb(),
			),
			service.NewSVMoveService(),
			presenter.NewGetMovesPresenter(),
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
