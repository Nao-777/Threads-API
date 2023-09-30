package repository

import (
	"threadsAPI/model"

	"gorm.io/gorm"
)

type IThreadRepository interface {
	CreateThread(thread *model.Thread) error
	GetThreadsByUserID(thread *[]model.Thread, userId string) error
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
	// if err := tr.db.First(threads, "user_id=?", userId).Error; err != nil {
	// 	return err
	// }
	if err := tr.db.Joins("User").Where("user_id=?", userId).Find(threads).Error; err != nil {
		return err
	}
	return nil
}

//threadデータ取得（取得件数）

//threadデータの削除

//threadデータの更新
