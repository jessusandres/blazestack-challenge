package routes

import (
	"blazestack.com/ms-incidents/cmd/controllers"
	"blazestack.com/ms-incidents/cmd/dto"
	"blazestack.com/ms-incidents/cmd/guards"
	"blazestack.com/ms-incidents/cmd/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router gin.IRouter) {
	authRouter := router.Group("/auth")

	authRouter.POST("/login", middlewares.ValidateJSON[dto.LoginDto](), controllers.Login)
	authRouter.GET("/profile", guards.AuthGuard(), controllers.Profile)
}
