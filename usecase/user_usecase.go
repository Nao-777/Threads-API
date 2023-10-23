package usecase

import (
	"fmt"
	"log"
	"os"
	"strings"
	"threadsAPI/model"
	"threadsAPI/repository"
	"threadsAPI/utility"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// ユーザusecaseのインターフェース
type IUserUsecase interface {
	GetUser(user model.User)(model.UserResponse, error)
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string,error)
	DeleteUser(user model.User)error
	UpdateUser(user model.User)error
	PostUserImg(user model.User)error
	//DeleteUserImg(user model.User)error
}

// ユーザusecaseの構造体
type userUsecase struct {
	ur repository.IUserRepository
	ut utility.IUtility
}

// ユーザusecaseのコンストラクタ
func NewUserUsecase(ur repository.IUserRepository,ut utility.IUtility) IUserUsecase {
	return &userUsecase{ur,ut}
}
//
func(uc *userUsecase)GetUser(user model.User)(model.UserResponse, error){
	
	if err:=uc.ur.GetUser(&user);err!=nil{
		return model.UserResponse{},err
	}
	imgBytes,err:=uc.ur.GetUserImg(&user)
	if err!=nil{
		return model.UserResponse{},err
	}
	imgB64:=uc.ut.ImgEndode(imgBytes)
	resUser:=model.UserResponse{
		ID: user.ID,
		LoginID: user.LoginID,
		Name: user.Name,
		ImageUrl: imgB64,
		CreatedAt: user.CreatedAt,
	}
	return resUser,nil
}

// サインアップ
func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	//パスワードをハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	//userIDを生成
	userUUId, err := uuid.NewRandom()
	if err != nil {
		return model.UserResponse{}, err
	}
	//ハイフンを除去、文字列にした値をユーザIDとして登録
	userId := strings.Replace(userUUId.String(), "-", "", -1)
	//user avaterをデコード
	if user.ImageUrl==""{
		img,err:=uu.ut.ImgFileEndode("sampleImg/noimage.jpeg")
		if err!=nil{
			return model.UserResponse{}, err
		}
		user.ImageUrl=img
	}
	uDec,err:=uu.ut.ImgDecode(user.ImageUrl)
	if err!=nil{
		log.Fatal(err)
	}
    remoteFileName:="avaterImg"
	remoteFilePath:=fmt.Sprintf("users/%s/avator/%s",userId,remoteFileName)
	//入力されたユーザの情報を登録
	newUser := model.User{
		ID:       userId,
		LoginID:  user.LoginID,
		Name:     user.Name,
		ImageUrl: remoteFilePath,
		Password: string(hash),
	}
	if err:=uu.ur.PostUserImg(&newUser,uDec);err!=nil{
		return model.UserResponse{},err
	}
	if err := uu.ur.InsertUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resposeData := model.UserResponse{
		ID:   newUser.ID,
		LoginID: newUser.LoginID,
		Name: newUser.Name,
		ImageUrl:newUser.ImageUrl,
		CreatedAt: newUser.CreatedAt,
	}
	return resposeData, nil
}

// サインイン
func (uc *userUsecase) Login(user model.User) (string,error) {
	loginId := user.LoginID
	storedUser := model.User{}
	if err := uc.ur.GetUserByLoginId(&storedUser, loginId); err != nil {
		return "",err
	}
	log.Printf("%+v", user)
	log.Printf("%+v", storedUser)
	//パスワードの認証
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "",err
	}
	privatekeyBytes,err:=os.ReadFile(os.Getenv("PATH_PRIVATEKEY"))
	if err!=nil{
		log.Fatal(err)
	}
	privatekey,err:=jwt.ParseRSAPrivateKeyFromPEM(privatekeyBytes)
	if err!=nil{
		log.Fatal(err)
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodRS512,jwt.MapClaims{
		"user_id":storedUser.ID,
		"exp":time.Now().Add(time.Hour*24).Unix(),
	})
	tokenString,err:=token.SignedString(privatekey)
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("tokenString:"+tokenString)
	return tokenString,nil
}
func(uu *userUsecase)DeleteUser(user model.User)error{
	loginId := user.LoginID
	storedUser := model.User{}
	if err := uu.ur.GetUserByLoginId(&storedUser, loginId); err != nil {
		return err
	}
	//パスワードの認証
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return err
	}
	if err:=uu.ur.DeleteUser(&user);err!=nil{
		return err
	}
	if storedUser.ImageUrl!=""{
		if err:=uu.ur.DeleteUserImg(&storedUser);err!=nil{
			return err
		}
	}
	return nil
}
func(uu *userUsecase)UpdateUser(user model.User)error{
	//パスワードをハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	//user avaterをデコード
	uDec,err:=uu.ut.ImgDecode(user.ImageUrl)
	if err!=nil{
		log.Fatal(err)
	}
    remoteFileName:="avaterImg"
	remoteFilePath:=fmt.Sprintf("users/%s/avator/%s",user.ID,remoteFileName)
	updateUser:=model.User{
		ID: user.ID,
		LoginID: user.LoginID,
		Name: user.Name,
		Password: string(hash),
		ImageUrl: remoteFilePath,
		UpdateAt: time.Now(),
	}
	if err:=uu.ur.PostUserImg(&updateUser,uDec);err!=nil{
		return err
	}
	if err:=uu.ur.UpDateUser(&updateUser);err!=nil{
		return err
	}
	return nil
}
func(uu *userUsecase)PostUserImg(user model.User)error{
	//user avaterをデコード
	uDec,err:=uu.ut.ImgDecode(user.ImageUrl)
	if err!=nil{
		log.Fatal(err)
	}
    remoteFileName:="avaterImg"
	remoteFilePath:=fmt.Sprintf("users/%s/avator/%s",user.ID,remoteFileName)
	updateUser:=model.User{
		ID: user.ID,
		ImageUrl: remoteFilePath,
		UpdateAt: time.Now(),
	}
	if err:=uu.ur.PostUserImg(&updateUser,uDec);err!=nil{
		return err
	}
	if err:=uu.ur.UpDateUser(&updateUser);err!=nil{
		return err
	}
	return nil
}
// func (uu *userUsecase)DeleteUserImg(user model.User)error{
// 	//アバターの削除をした場合は、noimage画像に切り替える
// 	img,err:=uu.ut.ImgFileEndode("sampleImg/noimage.jpeg")
// 	if err!=nil{
// 		return err
// 	}
// 	uDec,err:=uu.ut.ImgDecode(img)
// 	if err!=nil{
// 		log.Fatal(err)
// 	}
//     remoteFileName:="avaterImg"
// 	remoteFilePath:=fmt.Sprintf("users/%s/avator/%s",user.ID,remoteFileName)
// 	user.ImageUrl=remoteFilePath
// 	if err:=uu.ur.PostUserImg(&user,uDec);err!=nil{
// 		return err
// 	}
// 	if err := uu.ur.UpDateUser(&user); err != nil {
// 		return err
// 	}
// 	return nil
// }