// models/arrival.go
package models

import "gorm.io/gorm"

type Arrival struct {
	gorm.Model
	Date        string     `gorm:"not null"`
	Description string     `gorm:"not null"`
	SupplierID  uint       `gorm:"not null"`
	Supplier    *Supplier  `gorm:"foreignKey:SupplierID"`
	UserID      uint       `gorm:"not null"`
	User        *User
	TotalArrivalCost float64 `gorm:"-"`
	TotalQuantity    int     `gorm:"-"`
	ArrivalCompositions []ArrivalComposition `gorm:"foreignKey:ArrivalID"`
}

func (a *Arrival) CalculateTotals() {
	var totalCost float64
	var totalQuantity int
	for _, ac := range a.ArrivalCompositions {
		totalCost += ac.ItemTotal()
		totalQuantity += ac.Quantity
	}
	a.TotalArrivalCost = totalCost
	a.TotalQuantity = totalQuantity
}
