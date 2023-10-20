package model

import "time"

//ユーザ構造体
type User struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	LoginID   string    `json:"login_id" gorm:"unique"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	ImageUrl string `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

//レスポンスデータ用
type UserResponse struct {
	ID   string `json:"id"`
	LoginID string `json:"login_id"`
	Name string `json:"name"`
	ImageUrl string `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
