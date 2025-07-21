package entity

import "github.com/google/uuid"

// Tokens - структура для хранения токенов
type Tokens struct {
	AccessToken  string    `json:"accessToken" example:"eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxMjNlNDU2Ny1lODliLTEyZDMtYTQ1Ni00MjY2MTQxNzQwMDAiLCJhaWQiOiJhYzMzNDhiNy01MGQ3LTQ1NjMtYmE5NS02MzU5OWY5MWQ4NzEiLCJleHAiOjE3NTE5MTk5ODcsImlhdCI6MTc1MTkxOTM4N30.O2ZddFrqUbI33SZ3M5rHYDeJMaYzXrAgk13VP_xJIdIxgOAc-C4qtlGrSDDNqYDcvDWbSfNtJ2JmYm0vC0e8Ug"`
	RefreshToken string    `json:"-"`
	AccessId     uuid.UUID `json:"-"`
}
