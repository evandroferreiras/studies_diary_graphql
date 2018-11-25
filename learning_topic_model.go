package main

import (
	"github.com/jinzhu/gorm"
)

// LearningTopic is the base type for learningPath to be used by the db and gql
type LearningTopic struct {
	gorm.Model
	StudyID     uint
	Order       int32
	Description string
	Studies     []Study `gorm:"many2many:study_learningpath"`
}
