package main

import (
	"github.com/jinzhu/gorm"
)

// Reference is the base type for references to be used by the db and gql
type Reference struct {
	gorm.Model
	StudyID    uint
	URLLink    string
	CategoryID int
	Category   ReferenceCategory
}

// ReferenceCategory is the base type for a category to be used by the db and gql
type ReferenceCategory struct {
	gorm.Model
	Name string
	Pets []Reference
}
