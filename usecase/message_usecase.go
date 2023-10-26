package usecase

import (
	"fmt"
	"threadsAPI/constants"
	"threadsAPI/model"
	"threadsAPI/repository"
	"threadsAPI/utility"
)

type IMessageUsecase interface {
	CreateMessage(message *model.Message)error
	GetMessagesByThreadId(threadId string)([]model.ResMessage,error)
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
	msgId,err:=mu.ut.CreateUUID()
	if err!=nil{
		return err
	}
	message.Id=msgId
	if message.ImageUrl!=""{
		imgBytes,err:=mu.ut.ImgDecode(message.ImageUrl)
		if err!=nil{
			return err
		}
		remoteFileName:=constants.STORAGE_MESSAGE_IMG_NAME
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

func(mu *messageUsecase)GetMessagesByThreadId(threadId string)([]model.ResMessage,error){
	msg:=[]model.Message{}
	resMsgs:=[]model.ResMessage{}
	if err :=mu.mr.GetMessagesByThreadId(&msg,threadId);err!=nil{
		return []model.ResMessage{},err
	}
	for _,msg:=range msg{
		if msg.ImageUrl!=""{
			imgBytes,err:=mu.mr.GetMessageImg(&msg)
			if err!=nil{
				return []model.ResMessage{},err
			}
			img:=mu.ut.ImgEndode(imgBytes)
			msg.ImageUrl=img
		}
		resMsg:=model.ResMessage{
			Id: msg.Id,
			Name: msg.User.Name,
			AvatorImg: msg.User.ImageUrl,
			Message: msg.Message,
			ImageUrl: msg.ImageUrl,
			CreatedAt: msg.CreatedAt,
		}
		resMsgs=append(resMsgs, resMsg)
	}
	return resMsgs,nil
}
func(mu *messageUsecase)DeleteMessage(msgId string)error{
	storedMsg:=model.Message{}
	storedMsg.Id=msgId
	if err:=mu.mr.GetMessage(&storedMsg);err!=nil{
		return err
	}
	if storedMsg.ImageUrl!=""{
		if err :=mu.mr.DeleteMessageImg(&storedMsg);err!=nil{
			return err
		}
	}
	if err:=mu.mr.DeleteMessage(&storedMsg);err!=nil{
		return err
	}
	return nil
}
func(mu *messageUsecase)UpdateMessage(message *model.Message)error{
	if message.ImageUrl!=""{
		imgBytes,err:=mu.ut.ImgDecode(message.ImageUrl)
		if err!=nil{
			return err
		}
		remoteFileName:=constants.STORAGE_MESSAGE_IMG_NAME
		remoteFilePath:=fmt.Sprintf("messages/%s/main/%s",message.Id,remoteFileName)
		message.ImageUrl=remoteFilePath
		if err:=mu.mr.PostMessageImg(message,imgBytes);err!=nil{
			return err
		}
	}
	if err:=mu.mr.UpdateMessage(message);err !=nil{
		return err
	}
	return nil
}