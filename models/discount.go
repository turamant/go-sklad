// models/discount.go
package models

import "gorm.io/gorm"

type Discount struct {
	gorm.Model
	Name              string  `gorm:"not null"`
	Description       string  `gorm:"not null"`
	DiscountPercentage float64 `gorm:"not null"`
	UserID            uint    `gorm:"not null"`
	User              *User
}
