package router

import (
	"threadsAPI/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, tc controller.IThreadController) *echo.Echo {
	e := echo.New()
	//corsの設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		//アクセスを許可するフロントエンドドメインを追加
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET","POST","PUT","DELETE"},
	}))
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)

	t := e.Group("/threads")
	t.POST("", tc.CreateThread)
	t.GET("/:id", tc.GetThreadsByUserID)
	t.GET("", tc.GetThreads)
	return e
}
