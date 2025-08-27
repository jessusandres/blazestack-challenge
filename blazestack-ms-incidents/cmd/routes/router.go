package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BuildRoutes(router gin.IRouter) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "All right",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	apiRouter := router.Group("api/v1")

	apiRouter.GET("/fail", func(c *gin.Context) {
		println("fail!!")
		c.Error(errors.New("something went wrong"))

		return
	})

	IncidentsRouter(apiRouter)
	AuthRoutes(apiRouter)
}
