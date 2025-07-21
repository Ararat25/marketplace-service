package controller

import (
	"net/http"
)

// Logout godoc
// @Summary Деавторизация пользователя
// @Description Деактивирует текущую сессию пользователя по access токену в заголовке
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} StatusResponse "Пользователь успешно деавторизован"
// @Failure 400 {object} ErrorResponse "Ошибка деавторизации или отсутствует токен"
// @Security ApiKeyAuth
// @Router /api/logout [post]
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
