// models/expense.go
package models

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	Date        string     `gorm:"not null"`
	Description string     `gorm:"not null"`
	CustomerID  uint       `gorm:"not null"`
	Customer    *Customer  `gorm:"foreignKey:CustomerID"`
	UserID      uint       `gorm:"not null"`
	User        *User
	TotalExpenseCost float64 `gorm:"-"`
	TotalQuantity    int     `gorm:"-"`
	ExpenseCompositions []ExpenseComposition `gorm:"foreignKey:ExpenseID"`
}

func (e *Expense) CalculateTotals() {
	var totalCost float64
	var totalQuantity int
	for _, ec := range e.ExpenseCompositions {
		totalCost += ec.ItemTotal()
		totalQuantity += ec.Quantity
	}
	e.TotalExpenseCost = totalCost
	e.TotalQuantity = totalQuantity
}
