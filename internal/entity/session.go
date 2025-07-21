package entity

import (
	"github.com/google/uuid"
)

// Session - структура для таблицы session из базы данных
type Session struct {
	Id            int       `json:"id" gorm:"column:id;primaryKey"`
	UserId        uuid.UUID `json:"userId" gorm:"type:uuid;column:userId"`
	RefreshToken  string    `json:"refreshToken" gorm:"column:refreshToken"`
	AccessTokenID uuid.UUID `json:"accessTokenID" gorm:"type:uuid;column:accessTokenID"`
	User          User      `gorm:"foreignKey:UserId;references:Id"`
}
