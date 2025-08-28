package services

import (
	"fmt"

	"blazestack.com/ms-incidents/cmd/apperrors"
	"blazestack.com/ms-incidents/cmd/dto"
	"blazestack.com/ms-incidents/cmd/models"
	"blazestack.com/ms-incidents/cmd/types"
)

func CreateProduct(incidentToCreate dto.IncidentToCreate) (types.Incident, error) {

	var incident types.Incident

	if incidentToCreate.Title == "" {
		return incident, apperrors.NewBadRequestError("Title is required")
	}

	if types.IncidentTypes(incidentToCreate.IncidentType).IsValid() == false {
		return incident, apperrors.NewBadRequestError("Invalid or empty incident type")
	}

	newIncident := models.Incident{
		Title:        incidentToCreate.Title,
		Description:  incidentToCreate.Description,
		IncidentType: incidentToCreate.IncidentType,
		Location:     incidentToCreate.Location,
		Image:        incidentToCreate.Image,
	}

	ctx := models.DB.Create(&newIncident)

	if ctx.Error != nil {
		fmt.Println("Error creating incident:", ctx.Error.Error())

		return incident, apperrors.NewInternalServerError("Error creating incident")
	}

	incident = types.Incident{
		ID:           newIncident.ID,
		Title:        newIncident.Title,
		Description:  newIncident.Description,
		IncidentType: newIncident.IncidentType,
		Location:     newIncident.Location,
		Image:        newIncident.Image,
		CreationTime: newIncident.CreatedAt,
	}

	return incident, nil
}

func FetchAllIncidents() []types.Incident {
	var incidentsDB []models.Incident

	models.DB.Order("created_at DESC").Limit(10).Find(&incidentsDB)

	var incidents []types.Incident

	for _, incident := range incidentsDB {
		incidents = append(
			incidents,
			types.Incident{
				ID:           incident.ID,
				Title:        incident.Title,
				Description:  incident.Description,
				IncidentType: incident.IncidentType,
				Location:     incident.Location,
				CreationTime: incident.CreatedAt,
				Image:        incident.Image,
			},
		)
	}

	return incidents
}
