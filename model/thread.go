package model

import "time"

type Thread struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnUpdate:CASCADE,onDelete:CASCADE"`
	UserId    string    `json:"user_id" gorm:"not null"`
	Title     string    `json:"title"`
	Contents  string    `json:"contents"`
	ImageUrl string `json:"url"`
	StoragePath string `json:"storage_path"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}
type ResThread struct{
	ID string `json:"id"`
	UserName string `json:"user_name"`
	AvatorImg string `json:"avatorImg"`
	Title     string    `json:"title"`
	Contents  string    `json:"contents"`
	ImageUrl string `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}