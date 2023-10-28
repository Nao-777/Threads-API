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

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	LogOut(c echo.Context)error
	CsrfToken(c echo.Context)error
	DeleteUser(c echo.Context)error
	UpdateUser(c echo.Context)error
	GetUser(c echo.Context)error
}

type userController struct {
	uu usecase.IUserUsecase
	uv validation.IUserValidation
}

func NewUserController(uu usecase.IUserUsecase,uv validation.IUserValidation) IUserController {
	return &userController{uu,uv}
}
func(uc *userController)GetUser(c echo.Context)error{
	user:=model.User{}
	if err:=c.Bind(&user);err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	userRes,err:=uc.uu.GetUser(user)
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

// サインアップ
func (uc *userController) SignUp(c echo.Context) error {
	newUser := model.User{}
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err:=uc.uv.UserValidate(newUser,true);err!=nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.SignUp(newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

// サインイン
func (uc *userController) Login(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err:=uc.uv.UserValidate(user,false);err!=nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	token,err:=uc.uu.Login(user)
	if err !=nil{
		log.Fatal(err)
	}
	//log.Printf("%+v", user)
	log.Println(token)
	cookie:=new(http.Cookie)
	cookie.Name="token"
	cookie.Value=token
	cookie.Expires=time.Now().Add(24*time.Hour)
	cookie.Path="/"
	cookie.Domain="localhost"
	cookie.Secure=true
	cookie.HttpOnly=true
	cookie.SameSite=http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}
func(uc *userController)LogOut(c echo.Context)error{
	cookie:=new(http.Cookie)
	cookie.Name="token"
	cookie.Value=""
	cookie.Expires=time.Now()
	cookie.Path="/"
	cookie.Domain="localhost"
	cookie.Secure=true
	cookie.HttpOnly=true
	cookie.SameSite=http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}
func (uc *userController)CsrfToken(c echo.Context)error{
	token:=c.Get("csrf").(string)
	return c.JSON(http.StatusOK,echo.Map{
		"csrf_token":token,
	})
}
func (uc *userController)DeleteUser(c echo.Context)error{
	user:=model.User{}
	//userIDを受け取る
	if err:=c.Bind(&user);err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	if err:=uc.uu.DeleteUser(user);err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	cookie:=new(http.Cookie)
	cookie.Name="token"
	cookie.Value=""
	cookie.Expires=time.Now()
	cookie.Path="/"
	cookie.Domain="localhost"
	cookie.Secure=true
	cookie.HttpOnly=true
	cookie.SameSite=http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}
func(uc *userController)UpdateUser(c echo.Context)error{
	user:=model.User{}
	if err:=c.Bind(&user);err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	if err:=uc.uv.UserValidate(user,false);err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	if err:=uc.uu.UpdateUser(user);err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.NoContent(http.StatusOK)
}