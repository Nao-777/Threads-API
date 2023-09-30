package repository

import (
	"threadsAPI/model"

	"gorm.io/gorm"
)

type IThreadRepository interface {
	CreateThread(thread *model.Thread) error
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
func (tr *threadRepository) GetThreadsByUserID() error {
	return nil
}

//threadデータ取得（取得件数）

//threadデータの削除

//threadデータの更新
