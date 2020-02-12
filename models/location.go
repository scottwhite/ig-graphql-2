package models

import "gopkg.in/guregu/null.v3"

type Location struct {
	Id             string      `json:"id"`
	OrganizationId string      `json:"organization_id"`
	Name           string      `json:"name,string"`
	Street1        null.String `json:"street_1,string"`
	Street2        null.String `json:"street_2,string"`
	City           null.String `json:"city,string"`
	State          null.String `json:"state,string"`
	ZipCode        null.String `json:"zip_code,string"`
	GeoPoint       null.String `json:"geo_point"`
	Latitude       null.String `json:"latitude,string"`
	Longitude      null.String `json:"longitude,string"`
}
