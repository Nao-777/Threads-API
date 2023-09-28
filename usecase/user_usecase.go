package usecase

import (
	"threadsAPI/model"
	"threadsAPI/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// ユーザusecaseのインターフェース
type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
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
	userId, err := uuid.NewRandom()
	if err != nil {
		return model.UserResponse{}, err
	}
	//入力されたユーザの情報を登録
	newUser := model.User{
		ID:       userId.String(),
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
