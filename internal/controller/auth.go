package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type authReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

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
