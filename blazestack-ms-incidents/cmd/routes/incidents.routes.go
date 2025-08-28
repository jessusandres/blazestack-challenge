package routes

import (
	"blazestack.com/ms-incidents/cmd/controllers"
	"blazestack.com/ms-incidents/cmd/guards"
	"github.com/gin-gonic/gin"
)

func IncidentsRouter(router gin.IRouter) {
	productsRouter := router.Group("/incidents")

	productsRouter.Use(guards.AuthGuard())

	productsRouter.POST("", controllers.CreateIncident)

	productsRouter.GET("", controllers.FetchAllIncidents)

}
