package controller

import (
	"net/http"
	"threadsAPI/constants"
	"threadsAPI/controller/validation"
	"threadsAPI/model"
	"threadsAPI/usecase"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type IMessageController interface {
	CreateMessage(c echo.Context)error
	GetMessagesByThreadId(c echo.Context)error
	DeleteMessage(c echo.Context)error
	UpdateMessage(c echo.Context)error
}
type messageController struct {
	mu usecase.IMessageUsecase
	mv validation.IMessageValidation
}
func NewMessageController (mu usecase.IMessageUsecase,mv validation.IMessageValidation)IMessageController{
	return &messageController{mu,mv}
}
func (mc *messageController)CreateMessage(c echo.Context)error{
	user:=c.Get("user").(*jwt.Token)
	claims:=user.Claims.(jwt.MapClaims)
	userId:=claims[constants.JWT_USER_ID].(string)
	threadId:=c.Param(constants.PARAM_THREAD_ID)
	msg:=model.Message{
		UserId: userId,
		ThreadId: threadId,
	}
	if err:=c.Bind(&msg);err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	if err:=mc.mv.MessageValidate(msg);err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	if err:=mc.mu.CreateMessage(&msg);err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.NoContent(http.StatusOK)
}
func (mc *messageController)GetMessagesByThreadId(c echo.Context)error{
	threadId:=c.Param(constants.PARAM_THREAD_ID)
	msgs,err:=mc.mu.GetMessagesByThreadId(threadId)
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusBadRequest,msgs)
}
func (mc *messageController)DeleteMessage(c echo.Context)error{
	msg:=model.Message{}
	if err:=c.Bind(&msg);err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	if err:=mc.mu.DeleteMessage(msg.Id);err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.NoContent(http.StatusOK)
}
func(mc *messageController)UpdateMessage(c echo.Context)error{
	msg:=model.Message{}
	if err:=c.Bind(&msg);err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	if err:=mc.mv.MessageValidate(msg);err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	msg.UpdateAt=time.Now()
	if err:=mc.mu.UpdateMessage(&msg);err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.NoContent(http.StatusOK)
}