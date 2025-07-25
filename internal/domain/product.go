package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductCode string
	ProductName string
	Description string
	Price float64
	OrderDetail []OrderDetail `gorm:"foreignKey:ProductID;"`
}