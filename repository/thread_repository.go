package repository

import (
	"threadsAPI/model"

	"gorm.io/gorm"
)

type IThreadRepository interface {
	CreateThread(thread *model.Thread) error
	GetThreadsByUserID(thread *[]model.Thread, userId string) error
	GetThreadsLimitAndOffset(threads *[]model.Thread, limit int, offset int) error
	GetThreads(threads *[]model.Thread) error
}

type threadRepository struct {
	db *gorm.DB
}

func NewThreadRpository(db *gorm.DB) IThreadRepository {
	return &threadRepository{db}
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

//threadデータの更新
