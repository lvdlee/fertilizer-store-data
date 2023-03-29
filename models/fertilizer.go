package models

import "gorm.io/gorm"

type Fertilizer struct {
	gorm.Model
	name        string
	brand       string
	strength    int
	granularity float64
	density     float64
}
