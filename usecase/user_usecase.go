package usecase

import (
	"log"
	"os"
	"strings"
	"threadsAPI/model"
	"threadsAPI/repository"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// ユーザusecaseのインターフェース
type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string,error)
}

// ユーザusecaseの構造体
type userUsecase struct {
	ur repository.IUserRepository
}

// ユーザusecaseのコンストラクタ
func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

// サインアップ
func (uc *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
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
	//入力されたユーザの情報を登録
	newUser := model.User{
		ID:       userId,
		LoginID:  user.LoginID,
		Name:     user.Name,
		Password: string(hash),
	}
	if err := uc.ur.InsertUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resposeData := model.UserResponse{
		ID:   newUser.ID,
		Name: newUser.Name,
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
