package repository

import (
	"context"
	"threadsAPI/model"

	"cloud.google.com/go/storage"
	"gorm.io/gorm"
)

type IThreadRepository interface {
	CreateThread(thread *model.Thread) error
	GetThreadsByUserID(thread *[]model.Thread, userId string) error
	GetThreadsLimitAndOffset(threads *[]model.Thread, limit int, offset int) error
	GetThreads(threads *[]model.Thread) error
	DeleteThread(thread *model.Thread)error
	UpdateThread(thread *model.Thread)error
	PostThreadImg(thread *model.Thread,img []byte)error
}

type threadRepository struct {
	db *gorm.DB
	fbstorage *storage.BucketHandle
}

func NewThreadRpository(db *gorm.DB,fbstorage *storage.BucketHandle) IThreadRepository {
	return &threadRepository{db,fbstorage}
}

// threadデータの作成
func (tr *threadRepository) CreateThread(thread *model.Thread) error {
	if err := tr.db.Create(thread).Error; err != nil {
		return err
	}
	return nil
}

// threadデータの取得（ユーザIDで）
func (tr *threadRepository) GetThreadsByUserID(threads *[]model.Thread, userId string) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).Find(threads).Error; err != nil {
		return err
	}
	return nil
}

// threadデータ取得（取得件数）
func (tr *threadRepository) GetThreadsLimitAndOffset(threads *[]model.Thread, limit int, offset int) error {

	if err := tr.db.Joins("User").Offset(offset).Limit(limit).Find(threads).Error; err != nil {
		return err
	}
	return nil
}

// threadデータ取得（取得件数）
func (tr *threadRepository) GetThreads(threads *[]model.Thread) error {
	if err := tr.db.Joins("User").Find(threads).Error; err != nil {
		return err
	}
	return nil
}

//threadデータの削除
func (tr *threadRepository)DeleteThread(thread *model.Thread)error{
	if err:=tr.db.Delete(thread).Error;err!=nil{
		return err
	}
	return nil
}
//threadデータの更新
func (tr *threadRepository)UpdateThread(thread *model.Thread)error{
	if err:=tr.db.Updates(thread).Error;err!=nil{
		return err
	}
	return nil
}
func(tr *threadRepository)PostThreadImg(thread *model.Thread,img []byte)error{
	ctx:=context.Background()
	//storageで保管する画像の名前
	writer:=tr.fbstorage.Object("/test/").NewWriter(ctx)
	writer.ObjectAttrs.ContentType="image/jpg"
	writer.ObjectAttrs.CacheControl="no-cache"
	if _,err:=writer.Write(img);err!=nil{
		return err
	}
	defer writer.Close()
	return nil
}