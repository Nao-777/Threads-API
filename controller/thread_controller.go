package controller

import (
	"log"
	"net/http"
	"threadsAPI/controller/validation"
	"threadsAPI/model"
	"threadsAPI/usecase"
	"time"

	"github.com/labstack/echo/v4"
)

type IThreadController interface {
	CreateThread(c echo.Context) error
	//GetThreadsByUserID(c echo.Context) error
	GetThreads(c echo.Context) error
	DeleteThread(c echo.Context)error
	UpdateThread(c echo.Context)error
}
type threadController struct {
	tu usecase.IThreadUsecase
	tv validation.IThreadValidation
}

func NewThreadController(tu usecase.IThreadUsecase,tv validation.IThreadValidation) IThreadController {
	return &threadController{tu,tv}
}

// threadデータの作成
func (tc *threadController) CreateThread(c echo.Context) error {
	thread := model.Thread{}
	if err := c.Bind(&thread); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err:=tc.tv.ThreadValidate(thread);err!=nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := tc.tu.CreateThread(&thread); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

// threadデータの取得(ユーザID)
// func (tc *threadController) GetThreadsByUserID(c echo.Context) error {
// 	userId := c.Param("id")
// 	log.Printf("%s", userId)
// 	threads, err := tc.tu.GetThreadsByUserID(userId)
// 	if err != nil {
// 		return err
// 	}
// 	for _, v := range threads {
// 		log.Printf("%+v\n", v)
// 	}
// 	return c.JSON(http.StatusOK, threads)
// }

// threadデータの取得
func (tc *threadController) GetThreads(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	log.Printf("Get\n")
	log.Printf("%s:%s", limit, offset)
	//limitとoffsetの指定なし
	if limit == "" || offset == "" {
		threads, err := tc.tu.GetThreads()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, threads)
	}
	threads, err := tc.tu.GetThreadsLimitAndOffset(limit, offset)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, threads)
}
func(tc *threadController)DeleteThread(c echo.Context)error{
	thread:=model.Thread{}
	if err:=c.Bind(&thread);err!=nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err:=tc.tu.DeleteThread(thread);err!=nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
func (tc *threadController)UpdateThread(c echo.Context)error{
	thread:=model.Thread{}
	if err:=c.Bind(&thread);err!=nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err:=tc.tv.ThreadValidate(thread);err!=nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	thread.UpdateAt=time.Now()
	if err:=tc.tu.UpdateThread(thread);err!=nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
