package domain

import "gorm.io/gorm"

type OrderDetail struct {
	gorm.Model
	OrderID uint
	ProductID uint
	Amount int
}