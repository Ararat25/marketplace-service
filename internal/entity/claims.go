package entity

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// AccessTokenClaims - поля для access токена
type AccessTokenClaims struct {
	UserId        uuid.UUID `json:"userId"`
	AccessTokenID uuid.UUID `json:"aid"`
	jwt.RegisteredClaims
}

// RefreshTokenClaims - поля для refresh токена
type RefreshTokenClaims struct {
	UserId uuid.UUID `json:"userId"`
	jwt.RegisteredClaims
}
