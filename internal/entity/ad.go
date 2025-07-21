package entity

import (
	"github.com/google/uuid"
	"time"
)

type Ad struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title     string    `gorm:"size:100;not null" json:"title"`    // до 100 символов
	Content   string    `gorm:"size:1000;not null" json:"content"` // до 1000 символов
	ImageURL  string    `gorm:"size:255;not null" json:"imageUrl"` // до 255 символов
	Price     float64   `gorm:"check:price>=0" json:"price"`       // неотрицательная цена
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	User      User      `gorm:"foreignKey:UserID" json:"author"`
}
