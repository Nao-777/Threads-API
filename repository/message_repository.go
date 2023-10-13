package repository

import (
	"threadsAPI/model"

	"gorm.io/gorm"
)

type IMessageRepository interface {
	CreateMessage(message *model.Message)error
	GetMessages(message *[]model.Message,threadId string)error
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
func (mr *messageRepository)GetMessages(message *[]model.Message,threadId string)error{
	if err :=mr.db.Joins("Thread").Where("thread_id=?",threadId).Find(message).Error;err!=nil{
		return err
	}
	return nil
}