package dao

import "github.com/jmoiron/sqlx"

type Store struct {
	DB         *sqlx.DB
	Cis        *CisDAO
	Locations  *LocationsDAO
	Conductors *ConductorsDAO
	Schedules  *SchedulesDAO
}

func NewDAO(db *sqlx.DB) *Store {
	cis := NewCisDAO(db)
	locations := NewLocationDAO(db)
	conductors := NewConductorDAO(db)
	schedules := NewScheduleDAO(db)
	return &Store{db, cis, locations, conductors, schedules}
}
