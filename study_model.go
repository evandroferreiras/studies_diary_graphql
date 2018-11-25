package main

import (
	"github.com/jinzhu/gorm"
)

// Study is the base study model to be used throughout the app
type Study struct {
	gorm.Model
	ScopeDefinition   string
	SuccessDefinition string
	LearningPath      []LearningTopic
}

// StudyInput has everything needed to add an study
type StudyInput struct {
	ScopeDefinition   string
	SuccessDefinition string
}
