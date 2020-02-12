package models

import (
	"gopkg.in/guregu/null.v3"
)

type Ci struct {
	Id              string      `json:"id"`
	Name            string      `json:"name"`
	Comments        null.String `json:"comments,string"`
	StateID         string      `json:"state_id"`
	ModelID         null.String `json:"model_id,string"`
	LocationID      null.String `json:"location_id,string"`
	LocationFloorID null.String `json:"location_floor_id,string"`
	LocationRoomID  null.String `json:"location_room_id,string"`
	CabinetID       null.String `json:"cabinet_id,string"`
	OrganizationID  string      `json:"organization_id"`
	OsVersionID     null.String `json:"os_version_id,string"`
	LastChecked     null.String `json:"last_checked"`
	RecordSource    null.String `json:"record_source,string"`
	CollectorID     null.String `json:"collector_id,string"`
	CreatedOn       string      `json:"created_on"`
	BaselineAt      null.String `json:"baseline_at,string"`
	SSHPort         null.String `json:"ssh_port,string"`
	CiClass         string      `json:"ci_class"`
	IsCloud         bool        `json:"is_cloud"`
	ParentPath      null.String `json:"parent_path,string,omitempty"`
}

// type CiExtend struct {
// 	Ci
// 	LocationName     string `json:"location_name,string"`
// 	LocationStreet1  string `json:"location_street_1,string"`
// 	LocationStreet2  string `json:"location_street_2,string"`
// 	LocationCity     string `json:"location_city,string"`
// 	LocationState    string `json:"location_state,string"`
// 	LocationZipCode  string `json:"location_zip_code,string"`
// 	LocationLat      string `json:"location_latitude,string"`
// 	LocationLon      string `json:"location_longitude,string"`
// 	IPAddress        string `json:"ip_address,string"`
// 	OsVersionName    string `json:"os_version_name,string"`
// 	OsVersionVersion string `json:"os_version_version,string"`
// }
