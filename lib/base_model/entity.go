package base_model

import "gorm.io/gorm"

// BaseEntity is an interface that defines a method for loading associations
type BaseEntity interface {
	LoadWithAssociations(db *gorm.DB) *gorm.DB
}