// models/expensecomposition.go
package models

import "gorm.io/gorm"

type ExpenseComposition struct {
	gorm.Model
	ProductID   uint      `gorm:"not null"`
	Product     *Product  `gorm:"foreignKey:ProductID"`
	ExpenseID   uint      `gorm:"not null"`
	Expense     *Expense  `gorm:"foreignKey:ExpenseID"`
	Quantity    int       `gorm:"not null"`
	UserID      uint      `gorm:"not null"`
	User        *User
}

func (ec *ExpenseComposition) ItemTotal() float64 {
	discountedPrice := ec.Product.SellPrice * (1 - ec.Expense.Customer.Discount.DiscountPercentage/100)
	return discountedPrice * float64(ec.Quantity)
}
