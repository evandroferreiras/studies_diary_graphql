package main

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

// URL resolves the URLLink field for Study, it is all caps to avoid name clashes
func (u *Reference) URL(ctx context.Context) *string {
	return &u.URLLink
}

// CATEGORY resolves the category field for Study, it is all caps to avoid name clashes
func (u *Reference) CATEGORY(ctx context.Context) (*ReferenceCategory, error) {
	return db.getCategory(int32(u.CategoryID))
}

// ID resolves the category ID
func (u *ReferenceCategory) ID(ctx context.Context) *graphql.ID {
	return gqlIDP(u.Model.ID)
}

// NAME resolves the category name
func (u *ReferenceCategory) NAME(ctx context.Context) *string {
	return &u.Name
}
