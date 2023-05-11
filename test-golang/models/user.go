package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
}
