package entity

import "github.com/google/uuid"

// User - структура для таблицы user из базы данных
type User struct {
	Id       uuid.UUID `json:"id" gorm:"type:uuid;column:id"`
	Login    string    `json:"login" gorm:"column:login"`
	Password string    `json:"password" gorm:"column:password"`
	Sessions []Session `gorm:"foreignKey:UserId;references:Id"`
	Ads      []Ad      `gorm:"foreignKey:UserID;references:Id"`
}
