package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorResponse описывает структуру ответа с ошибкой
type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}

// StatusResponse описывает структуру ответа со статусом
type StatusResponse struct {
	Status string `json:"status" example:"ok"`
}

// sendSuccess отправляет успешный JSON-ответ с указанным статусом
func sendSuccess(w http.ResponseWriter, data any, statusCode int) {
	if data != nil {
		respBytes, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = w.Write(respBytes)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(statusCode)
}

// sendError отправляет JSON-ответ с ошибкой
func sendError(w http.ResponseWriter, errMsg string, statusCode int) {
	w.WriteHeader(statusCode)

	resp := ErrorResponse{Error: errMsg}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println("error encoding response:", err)
		http.Error(w, errMsg, statusCode)
	}
}
