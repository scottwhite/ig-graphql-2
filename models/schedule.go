package models

type ScheduleCore struct {
	Name           string  `json:"name,omitempty"`
	OrganizationID string  `json:"organization_id,string,omitempty"`
	Interval       int32   `json:"interval,omitempty"`
	ConductorID    *string `json:"conductor_id,string,omitempty"`
}

type Schedule struct {
	ID string `json:"id,string"`
	ScheduleCore
}

type ScheduleIn struct {
	ScheduleCore
}
