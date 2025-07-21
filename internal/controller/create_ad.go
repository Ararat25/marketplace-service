package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// createPostReq струкрура для запроса
type createPostReq struct {
	Title    string  `json:"title"`
	Content  string  `json:"content"`
	ImageURL string  `json:"imageUrl"`
	Price    float64 `json:"price"`
}

// createPostReq струкрура для ответа
type createPostResp struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	ImageURL  string    `json:"imageUrl"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}

// CreateAd godoc
// @Summary Создание нового объявления
// @Description Авторизованный пользователь может создать объявление
// @Tags ads
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ad body createPostReq true "Новое объявление"
// @Success 200 {object} createPostResp "Объявление успешно создано"
// @Failure 400 {object} ErrorResponse "Некорректные данные"
// @Failure 401 {object} ErrorResponse "Пользователь не авторизован"
// @Router /api/create [post]
func (h *Handler) CreateAd(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(uuid.UUID)
	if !ok {
		sendError(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var req createPostReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "invalid json", http.StatusBadRequest)
		return
	}

	ad, err := h.marketService.CreateAd(userID, req.Title, req.Content, req.ImageURL, req.Price)
	if err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	sendSuccess(w, createPostResp{
		ID:        ad.ID,
		Title:     ad.Title,
		Content:   ad.Content,
		ImageURL:  ad.ImageURL,
		Price:     ad.Price,
		CreatedAt: ad.CreatedAt,
	}, http.StatusOK)
}
