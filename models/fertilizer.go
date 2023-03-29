package models

import "gorm.io/gorm"

type Fertilizer struct {
	gorm.Model
	Name        string
	Brand       string
	Strength    int
	Granularity float64
	Density     float64
}
