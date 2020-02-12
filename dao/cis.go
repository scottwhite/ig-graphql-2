package dao

import (
	"context"
	"ig-graphql-2/models"

	"github.com/jmoiron/sqlx"
)

type CisDAO struct {
	db *sqlx.DB
}

func NewCisDAO(db *sqlx.DB) *CisDAO {
	return &CisDAO{db}
}

func (s *CisDAO) GetById(id string) (*models.Ci, error) {
	cis := make([]models.Ci, 0)
	err := s.db.Select(&cis, "select * from cis where $1", id)
	if err != nil {
		return nil, err
	}
	return &cis[0], nil
}

func (s *CisDAO) List(ctx context.Context) (*[]models.Ci, error) {
	cis := make([]models.Ci, 0)
	err := s.db.Select(&cis, "select * from cis LIMIT 300")
	if err != nil {
		return nil, err
	}
	return &cis, nil
}
