package passwordreset

import (
	bgorm "go-graphql-test/pkg/grom"
)

// PasswordReset domain
type PasswordReset struct {
	bgorm.BaseModel
	UserID uint   `gorm:"not null"`
	Token  string `gorm:"not null;unique_index"`
}
