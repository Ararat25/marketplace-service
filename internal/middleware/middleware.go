package middleware

import "github.com/Ararat25/api-marketplace/internal/model"

type Middleware struct {
	authService *model.AuthService
}

func NewMiddleware(authService *model.AuthService) *Middleware {
	return &Middleware{
		authService: authService,
	}
}
