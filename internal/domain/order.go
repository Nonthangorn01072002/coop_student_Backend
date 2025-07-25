package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ShipingAddress string
	Total float64
	UserID uint
	PaymentID uint
	Checkout_EmployeeID uint
	Payment Payment `gorm:"foreignKey:PaymentID"`
	OrderDetail []OrderDetail `gorm:"foreignKey:OrderID;"`
}