package dao

import (
	"context"
	"ig-graphql-2/models"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type ConductorsDAO struct {
	db *sqlx.DB
}

func NewConductorDAO(db *sqlx.DB) *ConductorsDAO {
	return &ConductorsDAO{db}
}

func (s *ConductorsDAO) GetById(id string) (*models.Conductor, error) {
	conductors := make([]models.Conductor, 0)
	err := s.db.Select(&conductors, "select * from conductors where id= $1", id)
	if err != nil {
		return nil, err
	}
	if len(conductors) == 0 {
		return &models.Conductor{}, nil
	}
	return &conductors[0], nil
}

func (s *ConductorsDAO) List(ctx context.Context) (*[]models.Conductor, error) {
	conductors := make([]models.Conductor, 0)
	err := s.db.Select(&conductors, "select * from conductors")
	if err != nil {
		return nil, err
	}
	if len(conductors) == 0 {
		return &[]models.Conductor{}, nil
	}
	return &conductors, nil
}

func (s *ConductorsDAO) Add(ctx context.Context, obj *map[string]interface{}) (*models.Conductor, error) {
	q := sq.Insert("conductors").SetMap(*obj).Suffix("RETURNING *").PlaceholderFormat(sq.Dollar)
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	log.Println("add conductor sql", query)

	res := make([]models.Conductor, 0)
	err = s.db.Select(&res, query, args...)
	if err != nil {
		return nil, err
	}
	return &res[0], nil
}

func (s *ConductorsDAO) Update(ctx context.Context, obj *map[string]interface{}) (*models.Conductor, error) {
	id := (*obj)["ID"]

	delete(*obj, "ID") //not updating this

	q := sq.Update("conductors").SetMap(*obj).Where(sq.Eq{"id": id}).Suffix("RETURNING *").PlaceholderFormat(sq.Dollar)
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	conductors := make([]models.Conductor, 0)

	err = s.db.Select(&conductors, query, args...)
	if err != nil {
		return nil, err
	}
	return &conductors[0], nil
}

func (s *ConductorsDAO) Delete(id string) (int64, error) {
	res, err := s.db.Exec("delete from conductors where id= $1", id)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}
