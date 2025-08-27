package models

type Incident struct {
	Id           uint   `json:"id" gorm:"primaryKey"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	IncidentType string `json:"incidentType" gorm:"column:incident_type"`
	Location     string `json:"location"`
	Image        string `json:"image"`
	Date         string `json:"date" gorm:"column:incident_date"`
	CreatedAt    string `json:"-" gorm:"-:save" gorm:"column:created_at"`
	UpdatedAt    string `json:"-" gorm:"-:save" gorm:"column:updated_at"`
}

func (m *Incident) TableName() string {
	return "incidents"
}
