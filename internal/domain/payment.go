package domain

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	PaymentType string
	PaymentStatus bool
	PaymentDate string
	Bank string
	TransferImage string
}