package repository

import (
	"threadsAPI/model"

	"gorm.io/gorm"
)

// ユーザリポジトリのインターフェース
type IUserRepository interface {
	GetUserByLoginId(user *model.User, loginId string) error
	InsertUser(user *model.User) error
}

// ユーザリポジトリの構造体
type userRepository struct {
	db *gorm.DB
}

// コンストラクタ
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
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
