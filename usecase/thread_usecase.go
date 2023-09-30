package usecase

import (
	"threadsAPI/model"
	"threadsAPI/repository"
)

type IThreadUsecase interface {
	CreateThread(thread *model.Thread) error
	GetThreadsByUserID(userId string) ([]model.Thread, error)
	GetThreads(limit int, offset int) ([]model.Thread, error)
}
type threadUsecase struct {
	tr repository.IThreadRepository
}

func NewThreadUsecase(tr repository.IThreadRepository) IThreadUsecase {
	return &threadUsecase{tr}
}

// threadデータの作成
func (tu *threadUsecase) CreateThread(thread *model.Thread) error {
	if err := tu.tr.CreateThread(thread); err != nil {
		return err
	}
	return nil
}

// threadデータの取得（userID）
func (tu *threadUsecase) GetThreadsByUserID(userId string) ([]model.Thread, error) {
	threads := []model.Thread{}
	if err := tu.tr.GetThreadsByUserID(&threads, userId); err != nil {
		return []model.Thread{}, err
	}
	return threads, nil
}

// threadデータの取得
func (tu *threadUsecase) GetThreads(limit int, offset int) ([]model.Thread, error) {
	threads := []model.Thread{}
	if err := tu.tr.GetThreads(&threads, limit, offset); err != nil {
		return []model.Thread{}, err
	}
	return threads, nil
}
