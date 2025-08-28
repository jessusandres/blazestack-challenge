package controllers

import (
	"fmt"
	"net/http"

	"blazestack.com/ms-incidents/cmd/apperrors"
	"blazestack.com/ms-incidents/cmd/dto"
	"blazestack.com/ms-incidents/cmd/helpers"
	"blazestack.com/ms-incidents/cmd/middlewares"
	"blazestack.com/ms-incidents/cmd/types"
	"github.com/gin-gonic/gin"
)

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	payload, success := middlewares.GetValidatedPayload[dto.LoginDto](c)

	if !success {
		return
	}

	fmt.Printf("Payload: %+v \n", payload)

	expectedUser := types.TokenClaims{
		Email:    "demo@example.com",
		Name:     "Jesús",
		LastName: "Cumpa",
		Lang:     "ES",
		Avatar:   "",
	}

	expectedPwd := "admin"

	if payload.Email != expectedUser.Email || payload.Password != expectedPwd {
		c.Error(
			apperrors.NewBadRequestError("wrong credentials"),
		)

		return
	}

	token, err := helpers.EncodeToken(expectedUser)

	if err != nil {
		c.Error(
			apperrors.NewInternalServerError("error generating token"),
		)

		return
	}
	
	response := types.Response{
		Data: gin.H{
			"token": token,
			"name":  expectedUser.Name + " " + expectedUser.LastName,
			"email": expectedUser.Email,
		},
	}

	c.JSON(
		http.StatusOK,
		response,
	)
}

func Profile(c *gin.Context) {
	user, exists := c.Get("user")

	if !exists {
		c.Error(
			apperrors.NewUnauthorizedError("user not found"),
		)
	}

	token, _ := c.Get("token")

	userClaims, ok := user.(types.TokenClaims)

	if !ok {
		c.Error(
			apperrors.NewInternalServerError("error parsing user claims"),
		)
	}

	fmt.Printf("User: %+v \n", userClaims)
	fmt.Printf("Token: %+v \n", token)

	response := types.Response{
		Data: gin.H{
			"token": token,
			"name":  userClaims.Name + " " + userClaims.LastName,
			"email": userClaims.Email,
		},
	}

	c.JSON(http.StatusOK, response)
}
