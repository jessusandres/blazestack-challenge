package controllers

import (
	"net/http"

	"blazestack.com/ms-incidents/cmd/dto"
	"blazestack.com/ms-incidents/cmd/helpers"
	"blazestack.com/ms-incidents/cmd/services"
	"blazestack.com/ms-incidents/cmd/types"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	productToCreate, ok := helpers.ValidateJsonPayload[dto.ProductToCreate](c)

	if !ok {
		return
	}

	product, err := services.CreateProduct(productToCreate)

	if err != nil {
		c.JSON(http.StatusBadRequest, types.Response{
			Data: nil,
		})

		return
	}

	response := types.Response{
		Data: types.ProductCreatedResponse{
			Product: product,
		},
	}

	c.JSON(
		http.StatusCreated,
		response,
	)
}
