package repository

import (
	"context"
	"io"
	"log"
	"os"
	"threadsAPI/model"

	"cloud.google.com/go/storage"
	"gorm.io/gorm"
)

// ユーザリポジトリのインターフェース
type IUserRepository interface {
	GetUserByLoginId(user *model.User, loginId string) error
	InsertUser(user *model.User) error
	DeleteUser(user *model.User)error
	UpDateUser(user *model.User)error
	PostUserImg(user *model.User)error
	GetUserImg(user *model.User)error
}

// ユーザリポジトリの構造体
type userRepository struct {
	db *gorm.DB
	fbstorage *storage.BucketHandle
}

// コンストラクタ
func NewUserRepository(db *gorm.DB,fbstorage *storage.BucketHandle) IUserRepository {
	return &userRepository{db,fbstorage}
}

// loginIdを使用してユーザ情報を取得する
func (ur *userRepository) GetUserByLoginId(user *model.User, loginId string) error {
	//loginIdを基にユーザを特定
	if err := ur.db.Where("login_id=?", loginId).First(&user).Error; err != nil {
		return err
	}
	return nil
}

// ユーザ情報の追加
func (ur *userRepository) InsertUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
func (ur *userRepository)DeleteUser(user *model.User)error{
	if err:=ur.db.Delete(user).Error;err!=nil{
		return err
	}
	return nil
}
func(ur *userRepository)UpDateUser(user *model.User)error{
	if err:=ur.db.Updates(user).Error;err!=nil{
		return err
	}
	return nil
}

func(ur *userRepository)PostUserImg(user *model.User)error{
	ctx:=context.Background()
	//storageで保管する画像の名前
	//仮
	remoteFileName:="asdfgh_"+user.ImageUrl
	writer:=ur.fbstorage.Object(remoteFileName).NewWriter(ctx)
	writer.ObjectAttrs.ContentType="image/jpg"
	writer.ObjectAttrs.CacheControl="no-cache"
	f,err:=os.Open(user.ImageUrl)
	if err!=nil{
		return err
	}
	if _,err:=io.Copy(writer,f);err!=nil{
		return err
	}
	defer f.Close()
	if err:=writer.Close();err!=nil{
		return err
	}
	return nil
}
func (ur *userRepository)GetUserImg(user *model.User)error{
	ctx:=context.Background()
	remoteFileName:="asdfgh_"+user.ImageUrl
	rc,err:=ur.fbstorage.Object(remoteFileName).NewReader(ctx)
	if err!=nil{
		return err
	}
	defer rc.Close()
	data,err:=io.ReadAll(rc)
	if err!=nil{
		return err
	}
	log.Printf("Download contents:%d\n",len(data))
	return nil
}