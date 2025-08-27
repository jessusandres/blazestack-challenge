package helpers

import (
	"fmt"
	"net/http"
	"os"

	"blazestack.com/ms-incidents/cmd/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ParseToken(c *gin.Context, tokenString string, reqUuid string) (types.TokenClaims, bool) {
	println("Trying to parse token: ", tokenString, " for request: ", reqUuid, "")

	claims := &types.TokenClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("algoritmo inesperado: %v", token.Header["alg"])
		}

		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("invalid algorithm: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		fmt.Printf("Error in JWT: %+v for request: %s \n", err, reqUuid)
		abortRequest(c, reqUuid)

		return *claims, false
	}

	if !token.Valid {
		println("Token invalid for request: ", reqUuid)
		abortRequest(c, reqUuid)

		return *claims, false
	}

	return *claims, true
}

func EncodeToken(claims types.TokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func abortRequest(c *gin.Context, uuid string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, types.ErrorResponse{
		Error: types.ErrorPayload{
			Message: "Wrong token provided",
			UUID:    uuid,
		},
	})
}
