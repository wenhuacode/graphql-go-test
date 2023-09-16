package global

import (
	"go-graphql-test/configs"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Config configs.Config
)
