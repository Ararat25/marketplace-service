package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// authReq структура для запроса
type authReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// Auth godoc
// @Summary Авторизация пользователя
// @Description Принимает логин и пароль, возвращает access и refresh токены в заголовках
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body authReq true "Данные для входа"
// @Success 200 {object} StatusResponse "Успешная авторизация"
// @Header 200 {string} X-Access-Token "JWT access токен"
// @Header 200 {string} X-Refresh-Token "JWT refresh токен"
// @Failure 400 {object} ErrorResponse "Неверный JSON или тело запроса"
// @Failure 409 {object} ErrorResponse "Неверный логин или пароль"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/auth [post]
func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req authReq
	err = json.Unmarshal(buf.Bytes(), &req)
	if err != nil {
		sendError(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	id, err := h.authService.VerifyUser(req.Login, req.Password)
	if err != nil {
		sendError(w, err.Error(), http.StatusConflict)
		return
	}

	tokens, err := h.authService.AuthUser(id)
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("X-Access-Token", tokens.AccessToken)
	w.Header().Set("X-Refresh-Token", tokens.RefreshToken)
	sendSuccess(w, StatusResponse{
		Status: "ok",
	}, http.StatusOK)
}
