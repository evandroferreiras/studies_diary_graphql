package main

func (db *DB) getCategory(id int32) (*ReferenceCategory, error) {
	var category ReferenceCategory
	err := db.DB.First(&category, id).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}
