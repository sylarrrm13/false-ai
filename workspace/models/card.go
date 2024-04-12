package models

import "utils"

type Card struct {
	ID         int             `gorm:"primary_key" json:"id"`
	CardTypeId int             `gorm:"not null" json:"card_type_id"`
	CardNo     string          `gorm:"not null" json:"card_no"`
	Use        *int            `gorm:"not null" json:"use"`
	UserId     *int            `gorm:"not null" json:"user_id"`
	CreatedAt  utils.LocalTime `gorm:"type:datetime not null" column:"created_at" `
	UpdatedAt  utils.LocalTime `gorm:"type:datetime not null" column:"updated_at"`
}

type CardUsage struct {
	ID        int             `gorm:"primary_key" json:"id"`
	Name      string          `gorm:"not null" json:"name"`
	CardNo    string          `gorm:"not null" json:"card_no"`
	Use       int             `gorm:"not null" json:"use"`
	UserName  string          `gorm:"not null" json:"username"`
	CreatedAt utils.LocalTime `gorm:"type:datetime not null;column:created_at" json:"created_at"`
	UpdatedAt utils.LocalTime `gorm:"type:datetime not null;column:updated_at" json:"updated_at"`
}

func (Card) TableName() string {
	return "card"
}

func (CardUsage) TableName() string {
	return "view_card_usage"
}
