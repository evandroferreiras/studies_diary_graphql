package main

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

// ID resolves the study ID
func (u *LearningTopic) ID(ctx context.Context) *graphql.ID {
	return gqlIDP(u.Model.ID)
}

// ORDER resolves the study order
func (u *LearningTopic) ORDER(ctx context.Context) int32 {
	return u.Order
}

// DESCRIPTION resolves the study description
func (u *LearningTopic) DESCRIPTION(ctx context.Context) *string {
	return &u.Description
}
