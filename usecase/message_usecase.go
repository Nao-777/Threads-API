package usecase

import (
	"strings"
	"threadsAPI/model"
	"threadsAPI/repository"

	"github.com/google/uuid"
)

type IMessageUsecase interface {
	CreateMessage(message *model.Message)error
	GetMessagesByThreadId(threadId string)([]model.Message,error)
	DeleteMessage(msgId string)error
}
type messageUsecase struct {
	mr repository.IMessageRepository
}

func NewMessageUsecase(mr repository.IMessageRepository)IMessageUsecase{
	return &messageUsecase{mr}
}
func (mu *messageUsecase)CreateMessage(message *model.Message)error{
	msgUUId,err:=uuid.NewRandom()
	if err !=nil{
		return err
	}
	msgId:=strings.Replace(msgUUId.String(),"-","",-1)
	message.Id=msgId
	if err:=mu.mr.CreateMessage(message);err!=nil{
		return err
	}
	return nil
}
func(mu *messageUsecase)GetMessagesByThreadId(threadId string)([]model.Message,error){
	msg:=[]model.Message{}
	if err :=mu.mr.GetMessagesByThreadId(&msg,threadId);err!=nil{
		return []model.Message{},err
	}
	return msg,nil
}
func(mu *messageUsecase)DeleteMessage(msgId string)error{
	msg:=model.Message{
		Id: msgId,
	}
	if err:=mu.mr.DeleteMessage(&msg);err!=nil{
		return err
	}
	return nil
}