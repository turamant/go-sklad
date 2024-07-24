// models/writeoff.go
package models

import "gorm.io/gorm"

type Writeoff struct {
	gorm.Model
	Date        string     `gorm:"not null"`
	Description string     `gorm:"not null"`
	ProductID   uint       `gorm:"not null"`
	Product     *Product   `gorm:"foreignKey:ProductID"`
	Quantity    int        `gorm:"not null"`
	UserID      uint       `gorm:"not null"`
	User        *User
	TotalWriteoffCost float64 `gorm:"-"`
}

func (w *Writeoff) CalculateTotalCost() {
	w.TotalWriteoffCost = w.Product.PurchasePrice * float64(w.Quantity)
}
