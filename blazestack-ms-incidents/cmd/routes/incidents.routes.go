package routes

import (
	"blazestack.com/ms-incidents/cmd/controllers"
	"blazestack.com/ms-incidents/cmd/guards"
	"blazestack.com/ms-incidents/cmd/services"
	"github.com/gin-gonic/gin"
)

func IncidentsRouter(router gin.IRouter) {
	productsRouter := router.Group("/incidents")

	productsRouter.POST("", guards.AuthGuard(), controllers.CreateProduct)

	productsRouter.GET("", guards.AuthGuard(), services.FetchAllProducts)

}
