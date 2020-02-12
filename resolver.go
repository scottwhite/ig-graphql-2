package main

import (
	"context"
	"fmt"
	"ig-graphql-2/dao"
	"ig-graphql-2/models"

	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
)

type Resolver struct {
	store *dao.Store
}

func NewResolver(dao *dao.Store) *Resolver {
	store = dao
	return &Resolver{store: dao}
}

var store *dao.Store

type CisResolver struct {
	models.Ci
}

func (cr *CisResolver) store() *dao.Store {
	return store
}

func (cr *CisResolver) ID(ctx context.Context) graphql.ID {
	r := graphql.ID(cr.Ci.Id)
	return r
}
func (cr *CisResolver) Name(ctx context.Context) string {
	return cr.Ci.Name
}
func (cr *CisResolver) OrganizationID(ctx context.Context) graphql.ID {
	r := graphql.ID(cr.Ci.OrganizationID)
	return r
}
func (cr *CisResolver) Comments(ctx context.Context) *string {
	return &cr.Ci.Comments.String
}
func (cr *CisResolver) StateID(ctx context.Context) graphql.ID {
	r := graphql.ID(cr.Ci.StateID)
	return r
}
func (cr *CisResolver) ModelId(ctx context.Context) *string {
	return &cr.Ci.ModelID.String
}
func (cr *CisResolver) Location(ctx context.Context) *LocationResolver {
	loc, err := store.Locations.GetById(cr.Ci.LocationID.String)
	if err != nil {
		fmt.Println("this sucks", err)
		return &LocationResolver{}
	}

	return &LocationResolver{Location: *loc}
}
func (cr *CisResolver) OsVersion(ctx context.Context) *string {
	return &cr.Ci.OsVersionID.String
}
func (cr *CisResolver) LastChecked(ctx context.Context) *string {
	return &cr.Ci.LastChecked.String
}
func (cr *CisResolver) RecordSource(ctx context.Context) *string {
	return &cr.Ci.RecordSource.String
}
func (cr *CisResolver) CreatedOn(ctx context.Context) string {
	return cr.Ci.CreatedOn
}
func (cr *CisResolver) BaseLineAt(ctx context.Context) *string {
	return &cr.Ci.BaselineAt.String
}
func (cr *CisResolver) SSHPort(ctx context.Context) *string {
	return &cr.Ci.SSHPort.String
}
func (cr *CisResolver) CiClass(ctx context.Context) string {
	return cr.Ci.CiClass
}
func (cr *CisResolver) IsCloud(ctx context.Context) bool {
	return cr.Ci.IsCloud
}
func (cr *CisResolver) CollectorID(ctx context.Context) *string {
	return &cr.Ci.CollectorID.String
}

//move to files this is getting HUGE
type LocationResolver struct {
	models.Location
}

func (cr *LocationResolver) ID(ctx context.Context) graphql.ID {
	r := graphql.ID(cr.Location.Id)
	return r
}
func (cr *LocationResolver) Name(ctx context.Context) string {
	return cr.Location.Name
}
func (cr *LocationResolver) OrganizationID(ctx context.Context) graphql.ID {
	r := graphql.ID(cr.Location.OrganizationId)
	return r
}
func (cr *LocationResolver) Street1(ctx context.Context) *string {
	return &cr.Location.Street1.String
}
func (cr *LocationResolver) Street2(ctx context.Context) *string {
	return &cr.Location.Street2.String
}
func (cr *LocationResolver) City(ctx context.Context) *string {
	return &cr.Location.City.String
}
func (cr *LocationResolver) ZipCode(ctx context.Context) *string {
	return &cr.Location.ZipCode.String
}
func (cr *LocationResolver) State(ctx context.Context) *string {
	return &cr.Location.State.String
}
func (cr *LocationResolver) GeoPoint(ctx context.Context) *string {
	return &cr.Location.GeoPoint.String
}
func (cr *LocationResolver) Latitude(ctx context.Context) *string {
	return &cr.Location.Latitude.String
}
func (cr *LocationResolver) Longitude(ctx context.Context) *string {
	return &cr.Location.Longitude.String
}

func (r *Resolver) GetCis(ctx context.Context) (*[]*CisResolver, error) {
	cis, err := r.store.Cis.List(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "GetCis")
	}
	cisr := make([]*CisResolver, 0)
	for _, ci := range *cis {
		cisr = append(cisr, &CisResolver{ci})
	}
	return &cisr, nil
}
