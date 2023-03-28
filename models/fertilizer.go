package models

import "gorm.io/gorm"

type Fertilizer struct {
	gorm.Model
	id          string
	name        string
	brand       string
	strength    int
	granularity float64
	density     float64
}
