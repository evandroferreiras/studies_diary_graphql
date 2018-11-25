package main

import (
	"github.com/jinzhu/gorm"
	// nolint: gotype
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB is the DB that will performs all operation
type DB struct {
	DB *gorm.DB
}

// NewDB returns a new DB connection
func newDB(path string) (*DB, error) {
	// connect to the example db, create if it doesn't exist
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// drop tables and all data, and recreate them fresh for this run
	db.DropTableIfExists(&Study{}, &LearningTopic{}, &Reference{}, &ReferenceCategory{})
	db.AutoMigrate(&Study{}, &LearningTopic{}, &Reference{}, &ReferenceCategory{})

	// put all the studies into the db
	for _, s := range studies {
		if err := db.Create(&s).Error; err != nil {
			return nil, err
		}
	}

	// put all the learningPath into the db
	for _, lp := range learningPath {
		if err := db.Create(&lp).Error; err != nil {
			return nil, err
		}
	}

	for _, c := range categories {
		if err := db.Create(&c).Error; err != nil {
			return nil, err
		}
	}

	for _, r := range references {
		if err := db.Create(&r).Error; err != nil {
			return nil, err
		}
	}

	return &DB{db}, nil
}

// TEST DATA TO BE PUT INTO THE DB
var studies = []Study{
	Study{ScopeDefinition: "Learn Go", SuccessDefinition: "Be able to create a console app in Go"},
	Study{ScopeDefinition: "Learn what is Single page app", SuccessDefinition: "Be able to explain to someone what is."},
	Study{ScopeDefinition: "Learn CSS", SuccessDefinition: "Be able to create a web site using bare css"},
}

// Since the db is torn down and created on every run, I know the above studies will have
// ID's 1, 2, 3
var learningPath = []LearningTopic{
	LearningTopic{Order: 1, Description: "Configure IDE", StudyID: 1},
	LearningTopic{Order: 2, Description: "Learn syntax", StudyID: 1},
	LearningTopic{Order: 3, Description: "Create console", StudyID: 1},

	LearningTopic{Order: 4, Description: "What is", StudyID: 2},
	LearningTopic{Order: 5, Description: "Advantages", StudyID: 2},
	LearningTopic{Order: 6, Description: "Disadvantages", StudyID: 2},

	LearningTopic{Order: 7, Description: "Learn about syntax", StudyID: 3},
	LearningTopic{Order: 8, Description: "Create a empty page", StudyID: 3},
	LearningTopic{Order: 9, Description: "Apply CSS", StudyID: 3},
}

// references to be put in the database
var references = []Reference{
	Reference{StudyID: 1, CategoryID: 1, URLLink: "https://some.site/"},
	Reference{StudyID: 1, CategoryID: 2, URLLink: "https://some.site/"},
	Reference{StudyID: 1, CategoryID: 3, URLLink: "https://some.site/"},
	Reference{StudyID: 2, CategoryID: 1, URLLink: "https://some.site/"},
	Reference{StudyID: 2, CategoryID: 2, URLLink: "https://some.site/"},
	Reference{StudyID: 2, CategoryID: 3, URLLink: "https://some.site/"},
	Reference{StudyID: 3, CategoryID: 1, URLLink: "https://some.site/"},
	Reference{StudyID: 3, CategoryID: 2, URLLink: "https://some.site/"},
	Reference{StudyID: 3, CategoryID: 3, URLLink: "https://some.site/"},
}

var categories = []ReferenceCategory{
	ReferenceCategory{Name: "Documentation"},
	ReferenceCategory{Name: "Books"},
	ReferenceCategory{Name: "Videos"},
}
