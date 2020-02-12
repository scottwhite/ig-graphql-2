package dao

import "github.com/jmoiron/sqlx"

type Store struct {
	DB        *sqlx.DB
	Cis       *CisDAO
	Locations *LocationsDAO
}

func NewDAO(db *sqlx.DB) *Store {
	cis := NewCisDAO(db)
	locations := NewLocationDAO(db)
	return &Store{db, cis, locations}
}
