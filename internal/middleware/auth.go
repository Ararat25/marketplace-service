package middleware

import (
	"context"
	"net/http"
)

// CheckAuth проверяет авторизован ли пользователь
func (m *Middleware) CheckAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Access-Token")
		if token == "" {
			http.Error(w, "missing access token", http.StatusUnauthorized)
			return
		}

		userID, err := m.authService.CheckAccess(token)
		if err != nil {
			http.Error(w, "invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Передаём userID в context
		ctx := context.WithValue(r.Context(), "userID", userID)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
