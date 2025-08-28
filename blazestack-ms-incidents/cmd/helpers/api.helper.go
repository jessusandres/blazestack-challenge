package helpers

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"blazestack.com/ms-incidents/cmd/types"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ExtractState(c *gin.Context) (types.AppState, bool) {
	reqState, exists := c.Get("state")

	if !exists {
		return types.AppState{}, false
	}

	state, ok := reqState.(types.AppState)

	return state, ok
}

func ValidateJsonPayload[T any](c *gin.Context) (T, bool) {
	var payload T

	if err := c.ShouldBindJSON(&payload); err != nil {
		var ve validator.ValidationErrors

		if errors.As(err, &ve) {
			out := make([]types.ValidationError, len(ve))

			for i, fe := range ve {
				out[i] = types.ValidationError{Field: fe.Field(), Message: GetErrorMsg(fe)}
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})

			return payload, false
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})

		return payload, false
	}

	return payload, true
}

func GetRequiredEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		panic(fmt.Sprintf("❌ Variable %s is required", key))
	}

	return value
}
