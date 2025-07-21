package controller

import (
	"net/http"
)

// refreshReq - структура для запроса на обновление токенов
type refreshReq struct {
	AccessToken string `json:"accessToken" example:"eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxMjNlNDU2Ny1lODliLTEyZDMtYTQ1Ni00MjY2MTQxNzQwMDAiLCJhaWQiOiJhYzMzNDhiNy01MGQ3LTQ1NjMtYmE5NS02MzU5OWY5MWQ4NzEiLCJleHAiOjE3NTE5MTk5ODcsImlhdCI6MTc1MTkxOTM4N30.O2ZddFrqUbI33SZ3M5rHYDeJMaYzXrAgk13VP_xJIdIxgOAc-C4qtlGrSDDNqYDcvDWbSfNtJ2JmYm0vC0e8Ug"`
}

// RefreshToken godoc
// @Summary Обновить пару токенов
// @Description Обновляет access и refresh токены, полученные из заголовков X-Access-Token и X-Refresh-Token. Проверяет User-Agent и IP, отправляет уведомление на webhook при смене IP.
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} entity.Tokens "Новая пара токенов"
// @Failure 400 {object} ErrorResponse "Ошибки валидации или отсутствуют токены"
// @Router /api/refresh [get]
func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	accessToken := r.Header.Get("X-Access-Token")
	refreshToken := r.Header.Get("X-Refresh-Token")

	if accessToken == "" {
		sendError(w, "missing access token in header", http.StatusBadRequest)
		return
	}
	if refreshToken == "" {
		sendError(w, "missing refresh token in header", http.StatusBadRequest)
		return
	}

	tokens, err := h.authService.RefreshToken(refreshToken, accessToken)
	if err != nil {
		sendError(w, "failed to refresh tokens: "+err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("X-Access-Token", tokens.AccessToken)
	w.Header().Set("X-Refresh-Token", tokens.RefreshToken)
	sendSuccess(w, StatusResponse{
		Status: "ok",
	}, http.StatusOK)
}
