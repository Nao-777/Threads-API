package repository

import (
	"context"
	"io"
	"log"
	"threadsAPI/model"

	"cloud.google.com/go/storage"
	"gorm.io/gorm"
)

// ユーザリポジトリのインターフェース
type IUserRepository interface {
	GetUser(user *model.User)error
	GetUserByLoginId(user *model.User, loginId string) error
	InsertUser(user *model.User) error
	DeleteUser(user *model.User)error
	UpDateUser(user *model.User)error
	PostUserImg(user *model.User,img []byte)error
	GetUserImg(user *model.User)([]byte,error)
	DeleteUserImg(user *model.User)error
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
func(ur *userRepository)GetUser(user *model.User)error{
	if err :=ur.db.First(&user).Error;err!=nil{
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

func(ur *userRepository)PostUserImg(user *model.User,img []byte)error{
	ctx:=context.Background()
	//storageで保管する画像の名前
	writer:=ur.fbstorage.Object(user.ImageUrl).NewWriter(ctx)
	writer.ObjectAttrs.ContentType="image/jpg"
	writer.ObjectAttrs.CacheControl="no-cache"
	if _,err:=writer.Write(img);err!=nil{
		return err
	}
	defer writer.Close()
	return nil
}
func (ur *userRepository)GetUserImg(user *model.User)([]byte,error){
	ctx:=context.Background()
	rc,err:=ur.fbstorage.Object(user.ImageUrl).NewReader(ctx)
	if err!=nil{
		return nil,err
	}
	defer rc.Close()
	data,err:=io.ReadAll(rc)
	if err!=nil{
		return nil,err
	}
	log.Printf("Download contents: %dbyte\n",len(data))
	return data,err
}
func(ur *userRepository)DeleteUserImg(user *model.User)error{
	ctx:=context.Background()
	if err:=ur.fbstorage.Object(user.ImageUrl).Delete(ctx);err!=nil{
		return err
	}
	return nil
}