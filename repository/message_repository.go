package repository

import (
	"threadsAPI/model"

	"gorm.io/gorm"
)

type IMessageRepository interface {
	CreateMessage(message *model.Message)error
	GetMessagesByThreadId(message *[]model.Message,threadId string)error
	DeleteMessage(message *model.Message)error
	UpdateMessage(message *model.Message)error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) IMessageRepository{
	return &messageRepository{db}
}
func (mr *messageRepository)CreateMessage(message *model.Message)error{
	if err:=mr.db.Create(message).Error;err!=nil{
		return err
	}
	return nil
}
func (mr *messageRepository)GetMessagesByThreadId(message *[]model.Message,threadId string)error{
	if err :=mr.db.Joins("Thread").Where("thread_id=?",threadId).Find(message).Error;err!=nil{
		return err
	}
	return nil
}
func (mr *messageRepository)DeleteMessage(message *model.Message)error{
	if err:=mr.db.Delete(message).Error;err!=nil{
		return err
	}
	return nil
}
func (mr *messageRepository)UpdateMessage(message *model.Message)error{
	if err:=mr.db.Updates(message).Error;err!=nil{
		return err
	}
	return nil
}