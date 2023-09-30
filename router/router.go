package router

import (
	"threadsAPI/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController, tc controller.IThreadController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)

	t := e.Group("/threads")
	t.POST("", tc.CreateThread)
	return e
}
