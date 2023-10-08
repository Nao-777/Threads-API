package router

import (
	"net/http"
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
		AllowHeaders:[]string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderXCSRFToken},
		AllowMethods: []string{"GET","POST","PUT","DELETE"},
		//cookieの送受信の許可
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:"/",
		CookieDomain:"localhost",
		CookieHTTPOnly: true,
		//CookieSameSite: http.SameSiteNoneMode,
		CookieSameSite:http.SameSiteDefaultMode,
		//CookieMaxAge:60,
	}))
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("logout",uc.LogOut)
	e.GET("/csrf",uc.CsrfToken)

	t := e.Group("/threads")
	t.POST("", tc.CreateThread)
	t.GET("/:id", tc.GetThreadsByUserID)
	t.GET("", tc.GetThreads)
	return e
}
