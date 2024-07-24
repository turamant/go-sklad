// models/customer.go
package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name        string     `gorm:"not null"`
	ContactInfo string     `gorm:"not null"`
	DiscountID  uint       `gorm:"null"`
	Discount    *Discount  `gorm:"foreignKey:DiscountID"`
	UserID      uint       `gorm:"not null"`
	User        *User
}
