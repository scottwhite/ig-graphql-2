package models

import "gopkg.in/guregu/null.v3"

type ConductorCore struct {
	OrganizationID string      `json:"organization_id,string"`
	Name           string      `json:"name"`
	Description    null.String `json:"description,string"`
	IpAddress      string      `json:"ip_address"`
	Port           int32       `json:"port"`
	Fingerprint    string      `json:"fingerprint"`
	LastPinged     null.Time   `json:"last_pinged"`
}

type Conductor struct {
	ID string `json:"id,string"`
	ConductorCore
}

type ConductorIn struct {
	ConductorCore
}
