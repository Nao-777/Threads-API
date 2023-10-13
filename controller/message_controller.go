package controller

import (
	"net/http"
	"threadsAPI/model"
	"threadsAPI/usecase"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type IMessageController interface {
	CreateMessage(c echo.Context)error
	GetMessagesByThreadId(c echo.Context)error
}
type messageController struct {
	mu usecase.IMessageUsecase
}
func NewMessageController (mu usecase.IMessageUsecase)IMessageController{
	return &messageController{mu}
}
func (mc *messageController)CreateMessage(c echo.Context)error{
	user:=c.Get("user").(*jwt.Token)
	claims:=user.Claims.(jwt.MapClaims)
	userId:=claims["user_id"].(string)
	msg:=model.Message{
		UserId: userId,
	}
	if err:=c.Bind(&msg);err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	if err:=mc.mu.CreateMessage(&msg);err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.NoContent(http.StatusOK)
}
func (mc *messageController)GetMessagesByThreadId(c echo.Context)error{
	threadId:=c.Param("threadId")
	msgs,err:=mc.mu.GetMessagesByThreadId(threadId)
	if err!=nil{
		return err
	}
	return c.JSON(http.StatusBadRequest,msgs)
}