package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	id       string
	username string
	password string
}
