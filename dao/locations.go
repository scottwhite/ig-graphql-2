package dao

import (
	"ig-graphql-2/models"

	"github.com/jmoiron/sqlx"
)

type LocationsDAO struct {
	db *sqlx.DB
}

func NewLocationDAO(db *sqlx.DB) *LocationsDAO {
	return &LocationsDAO{db}
}

func (s *LocationsDAO) GetById(id string) (*models.Location, error) {
	locations := make([]models.Location, 0)
	err := s.db.Select(&locations, "select * from locations where id= $1", id)
	if err != nil {
		return nil, err
	}
	if len(locations) == 0 {
		return &models.Location{}, nil
	}
	return &locations[0], nil
}
