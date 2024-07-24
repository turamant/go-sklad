// models/writeoffcomposition.go
package models

import "gorm.io/gorm"

type WriteoffComposition struct {
	gorm.Model
	WriteoffID uint      `gorm:"not null"`
	Writeoff   *Writeoff `gorm:"foreignKey:WriteoffID"`
	ProductID  uint      `gorm:"not null"`
	Product    *Product  `gorm:"foreignKey:ProductID"`
	Quantity   int       `gorm:"not null"`
	UserID     uint      `gorm:"not null"`
	User       *User
}

func (wc *WriteoffComposition) ItemTotal() float64 {
	return wc.Product.PurchasePrice * float64(wc.Quantity)
}
