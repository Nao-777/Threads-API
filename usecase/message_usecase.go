package usecase

import (
	"strings"
	"threadsAPI/model"
	"threadsAPI/repository"

	"github.com/google/uuid"
)

type IMessageUsecase interface {
	CreateMessage(message *model.Message)error
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