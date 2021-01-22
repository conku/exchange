package tests

import "github.com/conku/gorm"

type Product struct {
	gorm.Model
	Code       string
	Name       string
	Price      float64
	Tag        *string
	Category   Category
	CategoryID *uint
}

type Category struct {
	gorm.Model
	Name string
}
