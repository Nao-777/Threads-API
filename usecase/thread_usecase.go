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
	GetThreadsLimitAndOffset(limit string, offset string) ([]model.ResThread, error)
	GetThreads() ([]model.ResThread, error)
	DeleteThread(thread model.Thread)error
	UpdateThread(thread model.Thread)error
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
func (tu *threadUsecase) GetThreadsLimitAndOffset(limitParam string, offsetParam string) ([]model.ResThread, error) {
	threads := []model.Thread{}
	resThreads:=[]model.ResThread{}
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return []model.ResThread{}, err
	}
	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		return []model.ResThread{}, err
	}

	if err := tu.tr.GetThreadsLimitAndOffset(&threads, limit, offset); err != nil {
		return []model.ResThread{}, err
	}
	//res用の構造体に値を格納
	for _,thread:=range threads{
		resThread:=model.ResThread{
			ID: thread.ID,
			UserName: thread.User.Name,
			LoginID: thread.User.LoginID,
			Title: thread.Title,
			Contents: thread.Contents,
			CreatedAt: thread.CreatedAt,
		}
		resThreads = append(resThreads, resThread)
	}
	return resThreads, nil
}
func (tu *threadUsecase) GetThreads() ([]model.ResThread, error) {
	threads := []model.Thread{}
	resThreads:=[]model.ResThread{}
	if err := tu.tr.GetThreads(&threads); err != nil {
		return []model.ResThread{}, err
	}
	//res用の構造体に値を格納
	for _,thread:=range threads{
		resThread:=model.ResThread{
			ID: thread.ID,
			UserName: thread.User.Name,
			LoginID: thread.User.LoginID,
			Title: thread.Title,
			Contents: thread.Contents,
			CreatedAt: thread.CreatedAt,
		}
		resThreads = append(resThreads, resThread)
	}
	return resThreads, nil
}
func (tu *threadUsecase)DeleteThread(thread model.Thread)error{
	if err:=tu.tr.DeleteThread(&thread);err!=nil{
		return err
	}
	return nil
}
func(tu *threadUsecase)UpdateThread(thread model.Thread)error{
	if err:=tu.tr.UpdateThread(&thread);err!=nil{
		return err
	}
	return nil
}
