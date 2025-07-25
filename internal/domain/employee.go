package domain

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Fullname string
	Lastname string
	Age string
	WorkStartDate string
	Order []Order `gorm:"foreignKey:Checkout_EmployeeID"`
}