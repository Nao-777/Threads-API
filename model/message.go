package model

import "time"
type Message struct{
	Id string `json:"id" gorm:"primaryKey"`
	Thread Thread `json:"thread" gorm:"foreignKey:ThreadId; constraint:onDelete:CASCADE"`
	ThreadId string `json:"thread_id" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:onDelete:CASCADE"`
	UserId    string    `json:"user_id" gorm:"not null"`
	Message string `json:"message"`
	Url string `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}