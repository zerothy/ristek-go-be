package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Name     string
	Amount   int
	Types    string
	Category string
	Dates    string
	Note     string
}
