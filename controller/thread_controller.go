package controller

import (
	"net/http"
	"threadsAPI/model"
	"threadsAPI/usecase"

	"github.com/labstack/echo/v4"
)

type IThreadController interface {
	CreateThread(c echo.Context) error
}
type threadController struct {
	tu usecase.IThreadUsecase
}

func NewThreadController(tu usecase.IThreadUsecase) IThreadController {
	return &threadController{tu}
}

// threadデータの作成
func (tc *threadController) CreateThread(c echo.Context) error {
	thread := model.Thread{}
	if err := c.Bind(&thread); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := tc.tu.CreateThread(&thread); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return nil
}
