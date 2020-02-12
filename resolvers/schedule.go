package resolvers

import (
	"context"
	"fmt"
	"ig-graphql-2/dao"
	"ig-graphql-2/models"

	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
)

type ScheduleResolver struct {
	models.Schedule
}

func (cr *ScheduleResolver) store() *dao.Store {
	return store
}

func (cr *ScheduleResolver) ID(ctx context.Context) graphql.ID {
	r := graphql.ID(cr.Schedule.ID)
	return r
}
func (cr *ScheduleResolver) Name(ctx context.Context) string {
	return cr.Schedule.Name
}
func (cr *ScheduleResolver) OrganizationID(ctx context.Context) graphql.ID {
	r := graphql.ID(cr.Schedule.OrganizationID)
	return r
}
func (cr *ScheduleResolver) Interval(ctx context.Context) int32 {
	return cr.Schedule.Interval
}
func (cr *ScheduleResolver) Conductor(ctx context.Context) *ConductorResolver {
	conductor, err := store.Conductors.GetById(*cr.Schedule.ConductorID)
	if err != nil {
		fmt.Println("this sucks", err)
		return &ConductorResolver{}
	}

	return &ConductorResolver{*conductor}
}

func (r *Resolver) GetSchedule(ctx context.Context, args struct{ ID graphql.ID }) (*ScheduleResolver, error) {
	sched, err := r.store.Schedules.GetById(string(args.ID))
	if err != nil {
		return nil, errors.Wrap(err, "GetSchedule")
	}
	return &ScheduleResolver{*sched}, nil
}

func (r *Resolver) GetSchedules(ctx context.Context) (*[]*ScheduleResolver, error) {
	scheds, err := r.store.Schedules.List(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "GetSchedules")
	}
	sr := make([]*ScheduleResolver, 0)
	for _, s := range *scheds {
		sr = append(sr, &ScheduleResolver{s})
	}
	return &sr, nil
}

func (r *Resolver) CreateSchedule(ctx context.Context, args struct{ Schedule *models.ScheduleIn }) (*ScheduleResolver, error) {
	m := map[string]interface{}{
		"name":            args.Schedule.Name,
		"conductor_id":    &args.Schedule.ConductorID,
		"interval":        args.Schedule.Interval,
		"organization_id": args.Schedule.OrganizationID,
	}
	schedule, err := r.store.Schedules.Add(ctx, &m)
	if err != nil {
		return nil, errors.Wrap(err, "AddSchedule")
	}
	return &ScheduleResolver{*schedule}, nil
}

func (r *Resolver) UpdateSchedule(ctx context.Context, args struct{ Schedule *models.Schedule }) (*ScheduleResolver, error) {
	m := map[string]interface{}{
		"id":              args.Schedule.ID,
		"name":            args.Schedule.Name,
		"conductor_id":    args.Schedule.ConductorID,
		"interval":        args.Schedule.Interval,
		"organization_id": args.Schedule.OrganizationID,
	}

	schedule, err := r.store.Schedules.Update(ctx, &m)
	if err != nil {
		return nil, errors.Wrap(err, "UpdateSchedule")
	}
	return &ScheduleResolver{*schedule}, nil
}

func (r *Resolver) DeleteSchedule(ctx context.Context, args struct{ ID graphql.ID }) (*int32, error) {
	count, err := r.store.Schedules.Delete(string(args.ID))
	if err != nil {
		return nil, errors.Wrap(err, "DeleteSchedule")
	}
	c := int32(count)
	return &c, nil
}
