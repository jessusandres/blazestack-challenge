package models

import "time"

type Incident struct {
	ID           uint `gorm:"primaryKey"`
	Title        string
	Description  string
	IncidentType string `gorm:"column:incident_type"`
	Location     string
	Image        string
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime,column:updated_at"`
}

func (m *Incident) TableName() string {
	return "incidents"
}
