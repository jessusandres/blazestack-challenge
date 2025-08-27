package middlewares

import (
	"blazestack.com/ms-incidents/cmd/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func BuildState() gin.HandlerFunc {
	return func(c *gin.Context) {

		println("Running middleware for state:")

		newUuid, _ := uuid.NewUUID()
		reqUuid := newUuid.String()

		appState := types.AppState{
			Uuid: reqUuid,
			User: types.TokenClaims{},
		}

		c.Set("state", appState)

		c.Next()

	}
}
