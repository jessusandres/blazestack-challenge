package types

import "github.com/golang-jwt/jwt/v5"

type Response struct {
	Data interface{} `json:"data"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type AppState struct {
	Uuid string
	User TokenClaims
}

type ErrorPayload struct {
	Message string `json:"message"`
	UUID    string `json:"uuid"`
}

type ErrorResponse struct {
	Error ErrorPayload `json:"error"`
}

type TokenClaims struct {
	Email    string `json:"email"`
	Name     string `json:"role"`
	LastName string `json:"lastName"`
	Lang     string `json:"lang"`
	Avatar   string `json:"avatar"`
	jwt.RegisteredClaims
}
