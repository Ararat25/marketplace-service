package controller

import (
	"net/http"
)

// Logout godoc
// @Summary Деавторизация пользователя
// @Description Требует access токен, деактивирует текущую сессию.
// @Tags auth
// @Accept json
// @Produce json
// @Param accessToken body true "Access Token"
// @Success 200 {object} StatusResponse "Статус выполнения"
// @Failure 400 {object} ErrorResponse "Некорректный access токен или ошибка выхода"
// @Router /api/logout [get]
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	accessToken := r.Header.Get("X-Access-Token")
	if accessToken == "" {
		sendError(w, "missing access token in header", http.StatusBadRequest)
		return
	}

	_, err := h.authService.CheckAccess(accessToken)
	if err != nil {
		sendError(w, "cannot verify user", http.StatusBadRequest)
		return
	}

	err = h.authService.Logout(accessToken)
	if err != nil {
		sendError(w, "cannot logout user", http.StatusBadRequest)
		return
	}

	sendSuccess(w, StatusResponse{
		Status: "ok",
	}, http.StatusOK)
}
