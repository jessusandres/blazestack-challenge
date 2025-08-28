package routes

import (
	"errors"
	"net/http"

	"blazestack.com/ms-incidents/cmd/types"
	"github.com/gin-gonic/gin"
)

func BuildRoutes(router gin.IRouter) {

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	apiRouter := router.Group("api/v1")

	apiRouter.GET("/", func(c *gin.Context) {

		incidentTypes := []string{
			string(types.IncidentTypeEarthquake),
			string(types.IncidentTypeFire),
			string(types.IncidentTypeFlood),
		}

		data := map[string]any{
			"incidentTypes": incidentTypes,
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	apiRouter.GET("/fail", func(c *gin.Context) {
		println("fail!!")
		c.Error(errors.New("something went wrong"))

		return
	})

	IncidentsRouter(apiRouter)
	AuthRoutes(apiRouter)
}
