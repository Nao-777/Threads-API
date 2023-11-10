package repository

import (
	"context"
	"threadsAPI/model"
	"time"

	"cloud.google.com/go/storage"
	"gorm.io/gorm"
)

type IMessageRepository interface {
	CreateMessage(message *model.Message)error
	GetMessage(message *model.Message)error
	GetMessagesByThreadId(message *[]model.Message,threadId string)error
	DeleteMessage(message *model.Message)error
	UpdateMessage(message *model.Message)error
	PostMessageImg(message *model.Message,img []byte)error
	DeleteMessageImg(message *model.Message)error
	GetMessageImgUrl(message *model.Message)error
}

type messageRepository struct {
	db *gorm.DB
	fbstorage *storage.BucketHandle
}

func NewMessageRepository(db *gorm.DB,fbstorage *storage.BucketHandle) IMessageRepository{
	return &messageRepository{db,fbstorage}
}
func (mr *messageRepository)CreateMessage(message *model.Message)error{
	if err:=mr.db.Create(message).Error;err!=nil{
		return err
	}
	return nil
}
func(mr *messageRepository)GetMessage(message *model.Message)error{
	if err:=mr.db.First(message).Error;err!=nil{
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
func(mr *messageRepository)PostMessageImg(message *model.Message,img []byte)error{
	ctx:=context.Background()
	//storageで保管する画像の名前
	writer:=mr.fbstorage.Object(message.ImageUrl).NewWriter(ctx)
	writer.ObjectAttrs.ContentType="image/jpg"
	writer.ObjectAttrs.CacheControl="no-cache"
	if _,err:=writer.Write(img);err!=nil{
		return err
	}
	defer writer.Close()
	return nil
}
func(mr *messageRepository)DeleteMessageImg(message *model.Message)error{
	ctx:=context.Background()
	if err:=mr.fbstorage.Object(message.ImageUrl).Delete(ctx);err!=nil{
		return err
	}
	return nil
}
func(mr *messageRepository)GetMessageImgUrl(message *model.Message)error{
	//storageのパス
	object :=mr.fbstorage.Object(message.StoragePath)
	downloadURL,err:=mr.fbstorage.SignedURL(object.ObjectName(),&storage.SignedURLOptions{
		Expires: time.Now().AddDate(1, 0, 0),
		Method: "GET",
	})
	if err!=nil{
		return err
	}
	message.ImageUrl=downloadURL
	return nil
}