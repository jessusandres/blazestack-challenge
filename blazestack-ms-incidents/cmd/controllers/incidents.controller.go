package controllers

import (
	"net/http"

	"blazestack.com/ms-incidents/cmd/dto"
	"blazestack.com/ms-incidents/cmd/helpers"
	"blazestack.com/ms-incidents/cmd/services"
	"blazestack.com/ms-incidents/cmd/types"
	"github.com/gin-gonic/gin"
)

func CreateIncident(c *gin.Context) {
	imageEncoded, err := helpers.ExtractImageForm("image", c)

	if err != nil {
		c.Error(err)

		return
	}

	title := c.PostForm("title")
	incidentType := c.PostForm("incidentType")
	description := c.PostForm("description")
	location := c.PostForm("location")

	product, err := services.CreateProduct(
		dto.IncidentToCreate{
			Title:        title,
			Description:  description,
			IncidentType: incidentType,
			Location:     location,
			Image:        imageEncoded,
		},
	)

	if err != nil {
		c.Error(
			err)

		return
	}

	response := types.Response{
		Data: types.IncidentCreatedResponse{
			Incident: product,
		},
	}

	c.JSON(
		http.StatusCreated,
		response,
	)
}

func FetchAllIncidents(c *gin.Context) {
	incidents := services.FetchAllIncidents()

	c.JSON(
		http.StatusOK,
		types.Response{
			Data: types.IncidentsResponse{
				Incidents: incidents,
			},
		})
}
