package types

import "time"

type Incident struct {
	ID           uint      `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	IncidentType string    `json:"incidentType"`
	Location     string    `json:"location"`
	CreationTime time.Time `json:"creationTime"`
	Image        string    `json:"image"`
}

type IncidentCreatedResponse struct {
	Incident Incident `json:"incident"`
}

type IncidentsResponse struct {
	Incidents []Incident `json:"incidents"`
}

type IncidentTypes string

const (
	IncidentTypeFire       IncidentTypes = "fire"
	IncidentTypeEarthquake IncidentTypes = "earthquake"
	IncidentTypeFlood      IncidentTypes = "flood"
)

func (i IncidentTypes) IsValid() bool {
	switch i {
	case IncidentTypeFire, IncidentTypeEarthquake, IncidentTypeFlood:
		return true
	default:
		return false
	}
}
