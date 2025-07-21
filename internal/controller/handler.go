package controller

import (
	"github.com/Ararat25/api-marketplace/internal/model"
)

// Handler структура для обработчиков запросов
type Handler struct {
	authService   *model.AuthService
	marketService *model.MarketplaceService
}

// NewHandler создает новый объект Handler
func NewHandler(authService *model.AuthService, marketService *model.MarketplaceService) *Handler {
	return &Handler{
		authService:   authService,
		marketService: marketService,
	}
}
