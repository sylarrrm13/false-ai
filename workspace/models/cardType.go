package models

import "github.com/shopspring/decimal"

type CardType struct {
	ID       int             `gorm:"primary_key" json:"id"`
	Sort     int             `gorm:"not null" json:"sort"`
	Name     string          `gorm:"not null" json:"name"`
	Coin     int             `gorm:"not null" json:"coin"`
	Duration int             `gorm:"not null" json:"duration"`
	Price    decimal.Decimal `gorm:"not null" json:"price"`
	Discount *int            `gorm:"not null" json:"discount"`
}

type CardTypeNameList struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (CardType) TableName() string {
	return "card_type"
}
