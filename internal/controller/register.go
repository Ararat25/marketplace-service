package controller

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

// registerReq - структура запроса для регистрации
type registerReq struct {
	Login    string `json:"login" example:"user123"`
	Password string `json:"password" example:"securePassword"`
}

// registerResp - структура ответа после успешной регистрации
type registerResp struct {
	UserId uuid.UUID `json:"userId" example:"123e4567-e89b-12d3-a456-426614174000"`
	Login  string    `json:"login" example:"user123"`
}

// Register godoc
// @Summary Регистрация нового пользователя
// @Description Регистрирует нового пользователя с логином и паролем
// @Tags auth
// @Accept json
// @Produce json
// @Param input body registerReq true "Данные пользователя"
// @Success 200 {object} registerResp "Пользователь успешно зарегистрирован"
// @Failure 400 {object} ErrorResponse "Некорректный JSON"
// @Failure 409 {object} ErrorResponse "Пользователь уже существует"
// @Router /api/register [post]
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req registerReq
	err = json.Unmarshal(buf.Bytes(), &req)
	if err != nil {
		sendError(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	userId, err := h.authService.RegisterUser(req.Login, req.Password)
	if err != nil {
		sendError(w, err.Error(), http.StatusConflict)
		return
	}

	sendSuccess(w, registerResp{
		UserId: userId,
		Login:  req.Login,
	}, http.StatusOK)
}
