package usecase

import (
	"fmt"
	"log"
	"strconv"
	"threadsAPI/constants"
	"threadsAPI/model"
	"threadsAPI/repository"
	"threadsAPI/utility"
)

type IThreadUsecase interface {
	CreateThread(thread *model.Thread) error
	//GetThreadsByUserID(userId string) ([]model.Thread, error)
	GetThreadsLimitAndOffset(limit string, offset string) ([]model.ResThread, error)
	GetThreads() ([]model.ResThread, error)
	DeleteThread(thread model.Thread)error
	UpdateThread(thread model.Thread)error
}
type threadUsecase struct {
	tr repository.IThreadRepository
	ut utility.IUtility
}

func NewThreadUsecase(tr repository.IThreadRepository,ut utility.IUtility) IThreadUsecase {
	return &threadUsecase{tr,ut}
}

// threadデータの作成
func (tu *threadUsecase) CreateThread(thread *model.Thread) error {
	//スレッドIDの作成
	//userIDを生成
	threadId,err:=tu.ut.CreateUUID()
	if err!=nil{
		return err
	}
	thread.ID = threadId

	if thread.ImageUrl!=""{
		uDec,err:=tu.ut.ImgDecode(thread.ImageUrl)
		if err!=nil{
			log.Fatal(err)
		}
		remoteFileName:=constants.STORAGE_THREAD_IMG_NAME
		remoteFilePath:=fmt.Sprintf("threads/%s/main/%s",thread.ID,remoteFileName)
		thread.StoragePath=remoteFilePath
		if err:=tu.tr.PostThreadImg(thread,uDec);err!=nil{
			return err
		}
		if err:=tu.tr.GetThreadImgUrl(thread);err!=nil{
			return err
		}
	}

	if err := tu.tr.CreateThread(thread); err != nil {
		return err
	}
	return nil
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
	for _,thread :=range threads{
		resThread:=model.ResThread{
			ID: thread.ID,
			UserName: thread.User.Name,
			AvatorImg: thread.User.ImageUrl,
			Title: thread.Title,
			Contents: thread.Contents,
			ImageUrl: thread.ImageUrl,
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
	for _,thread :=range threads{
		resThread:=model.ResThread{
			ID: thread.ID,
			UserName: thread.User.Name,
			AvatorImg: thread.User.ImageUrl,
			Title: thread.Title,
			Contents: thread.Contents,
			ImageUrl: thread.ImageUrl,
			CreatedAt: thread.CreatedAt,
		}
		resThreads = append(resThreads, resThread)
	} 
	return resThreads, nil
}
func (tu *threadUsecase)DeleteThread(thread model.Thread)error{
	storedThread:=model.Thread{}
	storedThread.ID=thread.ID
	if err:=tu.tr.GetThread(&storedThread);err!=nil{
		return err
	}
	if storedThread.ImageUrl!=""{
		if err:=tu.tr.DeleteThreadImg(&storedThread);err!=nil{
			return err
		}
	}
	if err:=tu.tr.DeleteThread(&thread);err!=nil{
		return err
	}
	return nil
}
func(tu *threadUsecase)UpdateThread(thread model.Thread)error{
	if thread.ImageUrl!=""{
		imgBytes,err:=tu.ut.ImgDecode(thread.ImageUrl)
		if err !=nil{
			return err
		}
		remoteFileName:=constants.STORAGE_THREAD_IMG_NAME
		remoteFilePath:=fmt.Sprintf("threads/%s/main/%s",thread.ID,remoteFileName)
		thread.StoragePath=remoteFilePath
		if err:=tu.tr.PostThreadImg(&thread,imgBytes);err!=nil{
			return err
		}
	}
	if err:=tu.tr.UpdateThread(&thread);err!=nil{
		return err
	}
	return nil
}

