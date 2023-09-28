package usecase

import (
	"log"
	"strings"
	"threadsAPI/model"
	"threadsAPI/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// ユーザusecaseのインターフェース
type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) error
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
func (uc *userUsecase) Login(user model.User) error {
	//loginID
	loginId := user.LoginID
	//取得してきたユーザID
	storedUser := model.User{}
	//loginIDを基にユーザ情報を取得
	if err := uc.ur.GetUserByLoginId(&storedUser, loginId); err != nil {
		return err
	}
	log.Printf("%+v", user)
	log.Printf("%+v", storedUser)
	//パスワードの認証
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return err
	}
	return nil
}
