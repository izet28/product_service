package models

import (
	"gorm.io/gorm"
)

// Product mewakili struktur tabel produk di database dengan validasi
type Product struct {
	gorm.Model
	Name        string  `json:"name" validate:"required,min=3,max=100"`
	Description string  `json:"description" validate:"required,min=10"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Stock       int     `json:"stock" validate:"required,min=0"`
}
