// models/user.go
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Suppliers []Supplier
	Products  []Product
	Discounts []Discount
	Customers []Customer
	Arrivals  []Arrival
	Expenses  []Expense
	Returns   []Return
	Writeoffs []Writeoff
}
