package repository

import (
	"threadsAPI/model"

	"gorm.io/gorm"
)

type IMessageRepository interface {
	CreateMessage(message *model.Message)error
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