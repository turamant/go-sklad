// models/arrivalcomposition.go
package models

import "gorm.io/gorm"

type ArrivalComposition struct {
	gorm.Model
	ProductID     uint      `gorm:"not null"`
	Product       *Product  `gorm:"foreignKey:ProductID"`
	ArrivalID     uint      `gorm:"not null"`
	Arrival       *Arrival  `gorm:"foreignKey:ArrivalID"`
	PurchasePrice float64   `gorm:"not null"`
	Quantity      int       `gorm:"not null"`
	UserID        uint      `gorm:"not null"`
	User          *User
}

func (ac *ArrivalComposition) ItemTotal() float64 {
	return ac.PurchasePrice * float64(ac.Quantity)
}
