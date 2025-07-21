package controller

import (
	"github.com/google/uuid"
	"math"
	"net/http"
	"strconv"
)

// response структура для ответа
type response struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	ImageURL string    `json:"imageUrl"`
	Price    float64   `json:"price"`
	Author   string    `json:"author"`
	Mine     bool      `json:"mine,omitempty"`
}

// ListAd godoc
// @Summary Получение списка объявлений
// @Description Получает список объявлений с поддержкой пагинации, сортировки и фильтрации по цене. Если передан X-Access-Token, отмечает объявления пользователя как "mine".
// @Tags ads
// @Accept json
// @Produce json
// @Param page query int false "Номер страницы" default(1)
// @Param limit query int false "Количество на странице" default(10)
// @Param sort query string false "Поле сортировки: price, title, created_at"
// @Param order query string false "Порядок сортировки: asc или desc"
// @Param minPrice query number false "Минимальная цена"
// @Param maxPrice query number false "Максимальная цена"
// @Success 200 {array} response "Список объявлений"
// @Failure 500 {object} ErrorResponse "Ошибка загрузки объявлений"
// @Security ApiKeyAuth
// @Router /api/ads [post]
func (h *Handler) ListAd(w http.ResponseWriter, r *http.Request) {
	var ok bool
	var userID uuid.UUID

	token := r.Header.Get("X-Access-Token")
	if token != "" {
		id, err := h.authService.CheckAccess(token)
		if err == nil {
			ok = true
			userID = id
		}
	}

	page := parseIntQuery(r, "page", 1)
	limit := parseIntQuery(r, "limit", 10)
	sortBy := r.URL.Query().Get("sort")
	order := r.URL.Query().Get("order")
	minPrice := parseFloatQuery(r, "minPrice", 0)
	maxPrice := parseFloatQuery(r, "maxPrice", math.MaxFloat64)

	posts, err := h.marketService.ListAd(page, limit, sortBy, order, minPrice, maxPrice)
	if err != nil {
		sendError(w, "failed to load posts", http.StatusInternalServerError)
		return
	}

	var result []response
	if !ok {
		for _, p := range posts {
			result = append(result, response{
				ID:       p.ID,
				Title:    p.Title,
				Content:  p.Content,
				ImageURL: p.ImageURL,
				Price:    p.Price,
				Author:   p.User.Login,
			})
		}
	} else {
		for _, p := range posts {
			result = append(result, response{
				ID:       p.ID,
				Title:    p.Title,
				Content:  p.Content,
				ImageURL: p.ImageURL,
				Price:    p.Price,
				Author:   p.User.Login,
				Mine:     p.UserID == userID,
			})
		}
	}

	sendSuccess(w, result, http.StatusOK)
}

func parseIntQuery(r *http.Request, key string, def int) int {
	val := r.URL.Query().Get(key)
	if v, err := strconv.Atoi(val); err == nil {
		return v
	}
	return def
}

func parseFloatQuery(r *http.Request, key string, def float64) float64 {
	val := r.URL.Query().Get(key)
	if v, err := strconv.ParseFloat(val, 64); err == nil {
		return v
	}
	return def
}
