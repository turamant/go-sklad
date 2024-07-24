// models/return.go
package models

import "gorm.io/gorm"

type Return struct {
	gorm.Model
	Date        string     `gorm:"not null"`
	Description string     `gorm:"not null"`
	CustomerID  uint       `gorm:"not null"`
	Customer    *Customer  `gorm:"foreignKey:CustomerID"`
	UserID      uint       `gorm:"not null"`
	User        *User
	TotalReturnCost float64 `gorm:"-"`
	TotalQuantity   int     `gorm:"-"`
	ReturnCompositions []ReturnComposition `gorm:"foreignKey:ReturnID"`
}

func (r *Return) CalculateTotals() {
	var totalCost float64
	var totalQuantity int
	for _, rc := range r.ReturnCompositions {
		totalCost += rc.ItemTotal()
		totalQuantity += rc.Quantity
	}
	r.TotalReturnCost = totalCost
	r.TotalQuantity = totalQuantity
}
