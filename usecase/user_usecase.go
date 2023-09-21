package usecase

import (
	"threadsAPI/model"
	"threadsAPI/repository"
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

	//入力されたユーザの情報を登録
	newUser := model.User{
		ID:       "",
		LoginID:  user.LoginID,
		Name:     user.Name,
		Password: user.Password,
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
