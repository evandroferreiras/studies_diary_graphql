package main

import (
	"context"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
)

// Resolver is the root resolver
type Resolver struct{}

// GetStudy resolves the getStudy query
func (r *Resolver) GetStudy(ctx context.Context, args struct{ ID int32 }) (*Study, error) {
	return db.getStudy(ctx, int32(args.ID))
}

// AddStudy resolves the AddStudy mutation
func (r *Resolver) AddStudy(ctx context.Context, args struct{ Study StudyInput }) (*Study, error) {
	return db.addStudy(ctx, args.Study)
}

func gqlIDP(id uint) *graphql.ID {
	r := graphql.ID(fmt.Sprint(id))
	return &r
}
