// models/product.go
package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ArticleNumber       string  `gorm:"unique;not null"`
	Name                string  `gorm:"not null"`
	Description         string  `gorm:"not null"`
	PurchasePrice       float64 `gorm:"not null"`
	SellPrice           float64 `gorm:"not null"`
	Quantity            int     `gorm:"not null"`
	UserID              uint    `gorm:"not null"`
	User                *User
	TotalArrivalQuantity int     `gorm:"-"`
	TotalExpenseQuantity int     `gorm:"-"`
	BalanceQuantity      int     `gorm:"-"`
	BalanceValue         float64 `gorm:"-"`
	BalanceSell          float64 `gorm:"-"`
	TotalProfit          float64 `gorm:"-"`
	ArrivalCompositions []ArrivalComposition `gorm:"foreignKey:ArrivalID"`
	ExpenseCompositions []ExpenseComposition `gorm:"foreignKey:ExpenseID"`
}

func (p *Product) CalculateBalances() {
	p.TotalArrivalQuantity = p.calculateTotalArrivalQuantity()
	p.TotalExpenseQuantity = p.calculateTotalExpenseQuantity()
	p.BalanceQuantity = p.TotalArrivalQuantity - p.TotalExpenseQuantity
	p.BalanceValue = float64(p.BalanceQuantity) * p.PurchasePrice
	p.BalanceSell = float64(p.BalanceQuantity) * p.SellPrice
	p.TotalProfit = p.calculateTotalProfit()
}

func (p *Product) calculateTotalArrivalQuantity() int {
	var totalQuantity int
	for _, ac := range p.ArrivalCompositions {
		totalQuantity += ac.Quantity
	}
	return totalQuantity
}

func (p *Product) calculateTotalExpenseQuantity() int {
	var totalQuantity int
	for _, ec := range p.ExpenseCompositions {
		totalQuantity += ec.Quantity
	}
	return totalQuantity
}

func (p *Product) calculateTotalProfit() float64 {
	var totalRevenue float64
	for _, ec := range p.ExpenseCompositions {
		discountedPrice := ec.Product.SellPrice * (1 - ec.Expense.Customer.Discount.DiscountPercentage/100)
		totalRevenue += discountedPrice * float64(ec.Quantity)
	}

	totalCost := 0.0
	for _, ac := range p.ArrivalCompositions {
		totalCost += ac.PurchasePrice * float64(ac.Quantity)
	}
	totalCost = float64(p.TotalArrivalQuantity) * p.PurchasePrice

	return totalRevenue - totalCost
}
