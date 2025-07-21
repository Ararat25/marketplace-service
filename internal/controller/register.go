package controller

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

type registerReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type registerResp struct {
	UserId uuid.UUID `json:"userId"`
	Login  string    `json:"login"`
}

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
