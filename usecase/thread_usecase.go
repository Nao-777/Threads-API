package usecase

import (
	"strconv"
	"strings"
	"threadsAPI/model"
	"threadsAPI/repository"

	"github.com/google/uuid"
)

type IThreadUsecase interface {
	CreateThread(thread *model.Thread) error
	GetThreadsByUserID(userId string) ([]model.Thread, error)
	GetThreadsLimitAndOffset(limit string, offset string) ([]model.Thread, error)
	GetThreads() ([]model.Thread, error)
}
type threadUsecase struct {
	tr repository.IThreadRepository
}

func NewThreadUsecase(tr repository.IThreadRepository) IThreadUsecase {
	return &threadUsecase{tr}
}

// threadデータの作成
func (tu *threadUsecase) CreateThread(thread *model.Thread) error {
	//スレッドIDの作成
	//userIDを生成
	threadUUId, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	//ハイフンを除去、文字列にした値をユーザIDとして登録
	threadId := strings.Replace(threadUUId.String(), "-", "", -1)
	thread.ID = threadId
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
func (tu *threadUsecase) GetThreadsLimitAndOffset(limitParam string, offsetParam string) ([]model.Thread, error) {
	threads := []model.Thread{}
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return []model.Thread{}, err
	}
	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		return []model.Thread{}, err
	}

	if err := tu.tr.GetThreadsLimitAndOffset(&threads, limit, offset); err != nil {
		return []model.Thread{}, err
	}
	return threads, nil
}
func (tu *threadUsecase) GetThreads() ([]model.Thread, error) {
	threads := []model.Thread{}
	if err := tu.tr.GetThreads(&threads); err != nil {
		return []model.Thread{}, err
	}
	return threads, nil
}
