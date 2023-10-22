package usecase

import (
	"fmt"
	"strings"
	"threadsAPI/model"
	"threadsAPI/repository"
	"threadsAPI/utility"

	"github.com/google/uuid"
)

type IMessageUsecase interface {
	CreateMessage(message *model.Message)error
	GetMessagesByThreadId(threadId string)([]model.Message,error)
	DeleteMessage(msgId string)error
	UpdateMessage(message *model.Message)error
}
type messageUsecase struct {
	mr repository.IMessageRepository
	ut utility.IUtility
}

func NewMessageUsecase(mr repository.IMessageRepository,ut utility.IUtility)IMessageUsecase{
	return &messageUsecase{mr,ut}
}
func (mu *messageUsecase)CreateMessage(message *model.Message)error{
	msgUUId,err:=uuid.NewRandom()
	if err !=nil{
		return err
	}
	msgId:=strings.Replace(msgUUId.String(),"-","",-1)
	message.Id=msgId
	if message.ImageUrl!=""{
		imgBytes,err:=mu.ut.ImgDecode(message.ImageUrl)
		if err!=nil{
			return err
		}
		remoteFileName:="msgImg"
		remoteFilePath:=fmt.Sprintf("messages/%s/main/%s",message.Id,remoteFileName)
		message.ImageUrl=remoteFilePath
		if err:=mu.mr.PostMessageImg(message,imgBytes);err!=nil{
			return err
		}
	}
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
func(mu *messageUsecase)UpdateMessage(message *model.Message)error{
	if err:=mu.mr.UpdateMessage(message);err !=nil{
		return err
	}
	return nil
}