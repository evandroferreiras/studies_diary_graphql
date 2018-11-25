package main

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

// ID resolves the study ID
func (u *Study) ID(ctx context.Context) *graphql.ID {
	return gqlIDP(u.Model.ID)
}

// SCOPEDEFINITION resolves the ScopeDefinition field for Study, it is all caps to avoid name clashes
func (u *Study) SCOPEDEFINITION(ctx context.Context) *string {
	return &u.ScopeDefinition
}

// SUCCESSDEFINITION resolves the SuccessDefinition field for Study, it is all caps to avoid name clashes
func (u *Study) SUCCESSDEFINITION(ctx context.Context) *string {
	return &u.SuccessDefinition
}

// LEARNINGPATH resolves the LearningPath field for Study, it is all caps to avoid name clashes
func (u *Study) LEARNINGPATH(ctx context.Context) (*[]*LearningTopic, error) {
	return db.GetLearningPathForStudy(ctx, int32(u.Model.ID))
}

// REFERENCES resolver the references field for Study, it is all caps to avoid name clashes
func (u *Study) REFERENCES(ctx context.Context) (*[]*Reference, error) {
	return db.GetReferencesForStudy(ctx, int32(u.Model.ID))
}
