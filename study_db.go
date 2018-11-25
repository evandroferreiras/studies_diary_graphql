package main

import "context"

func (db *DB) getStudy(ctx context.Context, id int32) (*Study, error) {
	var study Study
	err := db.DB.First(&study, id).Error
	if err != nil {
		return nil, err
	}

	return &study, nil
}

// GetLearningPathForStudy gets LearningPath associated with the study
func (db *DB) GetLearningPathForStudy(ctx context.Context, id int32) (*[]*LearningTopic, error) {
	var lt []*LearningTopic
	err := db.DB.Where("study_id = ?", id).Find(&lt).Error
	if err != nil {
		return nil, err
	}
	return &lt, nil
}

// GetReferencesForStudy get References associated with the study
func (db *DB) GetReferencesForStudy(ctx context.Context, id int32) (*[]*Reference, error) {
	var ref []*Reference
	err := db.DB.Where("study_id = ?", id).Find(&ref).Error
	if err != nil {
		return nil, err
	}
	return &ref, nil
}

func (db *DB) addStudy(ctx context.Context, input StudyInput) (*Study, error) {

	study := Study{
		ScopeDefinition:   input.ScopeDefinition,
		SuccessDefinition: input.SuccessDefinition,
	}

	err := db.DB.Create(&study).Error
	if err != nil {
		return nil, err
	}

	return &study, nil
}
