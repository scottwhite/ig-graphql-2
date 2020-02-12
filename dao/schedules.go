package dao

import (
	"context"
	"ig-graphql-2/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type SchedulesDAO struct {
	db *sqlx.DB
}

func NewScheduleDAO(db *sqlx.DB) *SchedulesDAO {
	return &SchedulesDAO{db}
}

func (s *SchedulesDAO) GetById(id string) (*models.Schedule, error) {
	schedules := make([]models.Schedule, 0)
	err := s.db.Select(&schedules, "select * from schedules where id= $1", id)
	if err != nil {
		return nil, err
	}
	if len(schedules) == 0 {
		return &models.Schedule{}, nil
	}
	return &schedules[0], nil
}

func (s *SchedulesDAO) List(ctx context.Context) (*[]models.Schedule, error) {
	schedules := make([]models.Schedule, 0)
	err := s.db.Select(&schedules, "select * from schedules")
	if err != nil {
		return nil, err
	}
	if len(schedules) == 0 {
		return &[]models.Schedule{}, nil
	}
	return &schedules, nil
}

func (s *SchedulesDAO) Add(ctx context.Context, obj *map[string]interface{}) (*models.Schedule, error) {
	q := sq.Insert("schedules").SetMap(*obj).Suffix("RETURNING *").PlaceholderFormat(sq.Dollar)
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	res := make([]models.Schedule, 0)
	err = s.db.Select(&res, query, args...)
	if err != nil {
		return nil, err
	}
	return &res[0], nil
}

func (s *SchedulesDAO) Update(ctx context.Context, obj *map[string]interface{}) (*models.Schedule, error) {
	id := (*obj)["ID"]

	delete(*obj, "ID") //not updating this

	q := sq.Update("schedules").SetMap(*obj).Where(sq.Eq{"id": id}).Suffix("RETURNING *").PlaceholderFormat(sq.Dollar)
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	schedules := make([]models.Schedule, 0)

	err = s.db.Select(&schedules, query, args...)
	if err != nil {
		return nil, err
	}
	return &schedules[0], nil
}

func (s *SchedulesDAO) Delete(id string) (int64, error) {
	res, err := s.db.Exec("delete from schedules where id= $1", id)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}
