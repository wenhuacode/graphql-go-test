package user

import bgorm "go-graphql-test/pkg/grom"

// User domain model
type User struct {
	bgorm.BaseModel
	FirstName string `gorm:"size:255"`
	LastName  string `gorm:"size:255"`
	Email     string `gorm:"NOT NULL; UNIQUE_INDEX"`
	Password  string `gorm:"NOT NULL"`
	Role      string `gorm:"NOT_NULL;size:255;DEFAULT:'standard'"`
	Active    bool   `gorm:"NOT NULL; DEFAULT: true"`
	// Token    string `gorm:"NOT NULL; UNIQUE_INDEX"`
}
