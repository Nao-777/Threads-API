package controller

import (
	"log"
	"net/http"
	"threadsAPI/model"
	"threadsAPI/usecase"

	"github.com/labstack/echo/v4"
)

type IThreadController interface {
	CreateThread(c echo.Context) error
	GetThreadsByUserID(c echo.Context) error
	GetThreads(c echo.Context) error
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

// threadデータの取得(ユーザID)
func (tc *threadController) GetThreadsByUserID(c echo.Context) error {
	userId := c.Param("id")
	log.Printf("%s", userId)
	threads, err := tc.tu.GetThreadsByUserID(userId)
	if err != nil {
		return err
	}
	for _, v := range threads {
		log.Printf("%+v\n", v)
	}
	return c.JSON(http.StatusOK, threads)
}

// threadデータの取得
func (tc *threadController) GetThreads(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	log.Printf("%s:%s", limit, offset)
	//limitとoffsetの指定なし
	if limit == "" || offset == "" {
		//
	}
	threads, err := tc.tu.GetThreads(limit, offset)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, threads)
}
