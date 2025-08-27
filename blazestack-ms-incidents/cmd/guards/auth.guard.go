package guards

import (
	"net/http"
	"strings"

	"blazestack.com/ms-incidents/cmd/helpers"
	"blazestack.com/ms-incidents/cmd/types"
	"github.com/gin-gonic/gin"
)

func AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {

		uuid := ""
		state, ok := helpers.ExtractState(c)

		if ok {
			uuid = state.Uuid
		}

		tokenString := extractToken(c)

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, types.ErrorResponse{
				Error: types.ErrorPayload{
					Message: "Missing credentials",
					UUID:    uuid,
				},
			})

			return
		}

		claims, ok := helpers.ParseToken(c, tokenString, uuid)

		if !ok {
			return
		}

		c.Set("user", claims)
		c.Set("token", tokenString)

		c.Next()
	}
}

func extractToken(c *gin.Context) string {
	authorization := c.Request.Header.Get("Authorization")
	forwardedAuthorization := c.Request.Header.Get("x-forwarded-authorization")

	authHeader := authorization

	if authHeader == "" {
		authHeader = forwardedAuthorization
	}

	if authHeader == "" {
		return authHeader
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	return tokenString
}
