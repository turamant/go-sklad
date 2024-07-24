// models/returncomposition.go
package models

import "gorm.io/gorm"

type ReturnComposition struct {
	gorm.Model
	ProductID   uint      `gorm:"not null"`
	Product     *Product  `gorm:"foreignKey:ProductID"`
	ReturnID    uint      `gorm:"not null"`
	Return      *Return   `gorm:"foreignKey:ReturnID"`
	Quantity    int       `gorm:"not null"`
	UserID      uint      `gorm:"not null"`
	User        *User
}

func (rc *ReturnComposition) ItemTotal() float64 {
	return rc.Product.SellPrice * float64(rc.Quantity)
}
