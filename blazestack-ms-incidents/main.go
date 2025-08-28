package main

import (
	"fmt"
	"net/http"
	"time"

	"log"
	"os"

	"blazestack.com/ms-incidents/cmd/helpers"
	"blazestack.com/ms-incidents/cmd/middlewares"
	"blazestack.com/ms-incidents/cmd/models"
	"blazestack.com/ms-incidents/cmd/routes"
	"blazestack.com/ms-incidents/cmd/types"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("No .env file found")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	models.BuildConnection()

	router := gin.Default()

	router.Use(
		cors.New(cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}),
	)

	router.NoRoute(func(c *gin.Context) {
		routePath := c.Request.URL.Path
		routeMethod := c.Request.Method

		message := fmt.Sprintf("Route [%s] %s not found", routeMethod, routePath)
		uuid := ""

		state, ok := helpers.ExtractState(c)

		if ok {
			uuid = state.Uuid
		}

		c.JSON(http.StatusNotFound, types.ErrorResponse{
			Error: types.ErrorPayload{
				Message: message,
				UUID:    uuid,
			},
		})
	})

	router.Use(middlewares.BuildState())
	router.Use(middlewares.BuildErr())

	routes.BuildRoutes(router)

	router.Run("127.0.0.1:" + port)
}
