package dto

type IncidentToCreate struct {
	Title        string `json:"title" binding:"required"`
	Description  string `json:"description" binding:""`
	IncidentType string `json:"incidentType" binding:"required"`
	Location     string `json:"location" binding:""`
	Image        string `json:"image" binding:""`
}
