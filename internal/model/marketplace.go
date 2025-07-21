package model

import (
	"errors"
	"fmt"
	"github.com/Ararat25/api-marketplace/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// MarketplaceService - структура для маркетплейса
type MarketplaceService struct {
	Storage *gorm.DB
}

// NewMarketplaceService возвращает новый объект структуры MarketplaceService
func NewMarketplaceService(storage *gorm.DB) *MarketplaceService {
	return &MarketplaceService{
		Storage: storage,
	}
}

func (s *MarketplaceService) CreateAd(userID uuid.UUID, title, content, imageURL string, price float64) (*entity.Ad, error) {
	if len(title) < 3 || len(title) > 100 || len(content) < 10 || len(content) > 1000 || price < 0 || len(imageURL) > 255 {
		return nil, errors.New("validation failed")
	}

	ad := &entity.Ad{
		ID:        uuid.New(),
		Title:     title,
		Content:   content,
		ImageURL:  imageURL,
		Price:     price,
		UserID:    userID,
		CreatedAt: time.Now(),
	}

	err := s.Storage.Create(ad).Error
	if err != nil {
		return nil, err
	}

	return ad, nil
}

func (s *MarketplaceService) ListAd(page, limit int, sortBy, order string, minPrice, maxPrice float64) ([]entity.Ad, error) {
	if sortBy != "price" {
		sortBy = "created_at"
	}
	if order != "asc" {
		order = "desc"
	}

	var posts []entity.Ad
	err := s.Storage.Preload("User").
		Where("price BETWEEN ? AND ?", minPrice, maxPrice).
		Order(fmt.Sprintf("%s %s", sortBy, order)).
		Limit(limit).
		Offset((page - 1) * limit).
		Find(&posts).Error

	return posts, err
}
