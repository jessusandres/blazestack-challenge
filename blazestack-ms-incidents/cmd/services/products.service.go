package services

import (
	"fmt"
	"net/http"

	"blazestack.com/ms-incidents/cmd/dto"
	"blazestack.com/ms-incidents/cmd/models"
	"blazestack.com/ms-incidents/cmd/types"
	"github.com/gin-gonic/gin"
)

func CreateProduct(productToCreate dto.ProductToCreate) (types.Product, error) {

	fmt.Printf("data: %+v \n", productToCreate)

	product := types.Product{
		ID:      1,
		Name:    productToCreate.Name,
		Sku:     productToCreate.Sku,
		Visible: false,
	}

	return product, nil
}

type ProductsResponse struct {
	Incidents []models.Incident `json:"incidents"`
}

func FetchAllProducts(c *gin.Context) {
	var incidents []models.Incident

	models.DB.Find(&incidents)

	fmt.Printf("incidents: %+v \n", incidents)

	c.JSON(
		http.StatusOK,
		types.Response{
			Data: ProductsResponse{
				Incidents: incidents,
			},
		})
}
