// models/supplier.go
package models

import (
	"gorm.io/gorm"
)

type Supplier struct {
	gorm.Model
	Name        string `gorm:"not null"`
	ContactInfo string `gorm:"not null"`
	UserID      uint   `gorm:"not null"`
	User        *User
	Arrivals    []Arrival `gorm:"foreignKey:SupplierID"`
}

