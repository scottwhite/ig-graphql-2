package resolvers

import (
	"context"
	"ig-graphql-2/dao"
	"ig-graphql-2/models"

	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
)

type ConductorResolver struct {
	models.Conductor
}

func (cr *ConductorResolver) store() *dao.Store {
	return store
}

func (cr *ConductorResolver) ID(ctx context.Context) graphql.ID {
	r := graphql.ID(cr.Conductor.ID)
	return r
}
func (cr *ConductorResolver) Name(ctx context.Context) string {
	return cr.Conductor.Name
}
func (cr *ConductorResolver) Description(ctx context.Context) *string {
	return &cr.Conductor.Description.String
}
func (cr *ConductorResolver) OrganizationID(ctx context.Context) graphql.ID {
	r := graphql.ID(cr.Conductor.OrganizationID)
	return r
}
func (cr *ConductorResolver) IpAddress(ctx context.Context) string {
	return cr.Conductor.IpAddress
}
func (cr *ConductorResolver) Port(ctx context.Context) int32 {
	return cr.Conductor.Port
}
func (cr *ConductorResolver) Fingerprint(ctx context.Context) string {
	return cr.Conductor.Fingerprint
}
func (cr *ConductorResolver) LastPinged(ctx context.Context) *string {
	r := cr.Conductor.LastPinged.Time.String()
	return &r
}

func (r *Resolver) GetConductor(ctx context.Context, args struct{ ID graphql.ID }) (*ConductorResolver, error) {
	cs, err := r.store.Conductors.GetById(string(args.ID))
	if err != nil {
		return nil, errors.Wrap(err, "GetConductor")
	}
	return &ConductorResolver{*cs}, nil
}

func (r *Resolver) GetConductors(ctx context.Context) (*[]*ConductorResolver, error) {
	cs, err := r.store.Conductors.List(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "GetConductors")
	}
	csr := make([]*ConductorResolver, 0)
	for _, c := range *cs {
		csr = append(csr, &ConductorResolver{c})
	}
	return &csr, nil
}

func (r *Resolver) CreateConductor(ctx context.Context, args struct{ Conductor *models.ConductorIn }) (*ConductorResolver, error) {

	m := map[string]interface{}{
		"name":            args.Conductor.Name,
		"description":     args.Conductor.Description,
		"fingerprint":     args.Conductor.Fingerprint,
		"last_pinged":     args.Conductor.LastPinged,
		"organization_id": args.Conductor.OrganizationID,
		"port":            args.Conductor.Port,
		"ip_address":      args.Conductor.IpAddress,
	}

	conductor, err := r.store.Conductors.Add(ctx, &m)
	if err != nil {
		return nil, errors.Wrap(err, "AddConductor")
	}
	return &ConductorResolver{*conductor}, nil
}

func (r *Resolver) UpdateConductor(ctx context.Context, args struct{ Conductor *models.Conductor }) (*ConductorResolver, error) {
	m := map[string]interface{}{
		"id":              args.Conductor.ID,
		"name":            args.Conductor.Name,
		"description":     args.Conductor.Description,
		"fingerprint":     args.Conductor.Fingerprint,
		"last_pinged":     args.Conductor.LastPinged,
		"organization_id": args.Conductor.OrganizationID,
		"port":            args.Conductor.Port,
		"ip_address":      args.Conductor.IpAddress,
	}

	conductor, err := r.store.Conductors.Update(ctx, &m)
	if err != nil {
		return nil, errors.Wrap(err, "UpdateConductor")
	}
	return &ConductorResolver{*conductor}, nil
}

func (r *Resolver) DeleteConductor(ctx context.Context, args struct{ ID graphql.ID }) (*int32, error) {
	count, err := r.store.Conductors.Delete(string(args.ID))
	if err != nil {
		return nil, errors.Wrap(err, "DeleteConductor")
	}
	c := int32(count)
	return &c, nil
}
