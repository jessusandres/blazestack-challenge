package routes

import (
	"blazestack.com/ms-incidents/cmd/controllers"
	"blazestack.com/ms-incidents/cmd/guards"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router gin.IRouter) {
	authRouter := router.Group("/auth")

	authRouter.POST("/login", controllers.Login)
	authRouter.GET("/profile", guards.AuthGuard(), controllers.Profile)
}
