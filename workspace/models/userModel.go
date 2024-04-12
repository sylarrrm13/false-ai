package models

import "utils"

type User struct {
	Id        int             `gorm:"primary_key;column:id" `
	Username  string          `gorm:"not null;unique" column:"username"`
	Password  string          `gorm:"not null" column:"password"`
	Email     string          `gorm:"unique;column:email"`
	Phone     string          `column:"phone"`
	Ref       string          `column:"ref"`  //预留字段可能用于存储微信相关信息
	Ref2      string          `column:"ref2"` //预留字段可能用于存储支付相关信息
	Ref3      string          `column:"ref3"` //预留字段
	Status    int             `gorm:"not null;column:status" `
	CreatedAt utils.LocalTime `gorm:"type:datetime not null;column:created_at"  `
	UpdatedAt utils.LocalTime `gorm:"type:datetime not null; column:updated_at"`
}

type UserBillInfo struct {
	Id           int              `gorm:"primary_key;column:id" `
	Userid       int              `gorm:"not null;column:userid" `
	Coins        int              `column:"coins"` //剩余积分
	UsedCoins    int              `column:"used_coins"`
	CoinsExpired *utils.LocalTime `gorm:"type:datetime;column:coins_expired"`
	Expired      *utils.LocalTime `gorm:"type:datetime;column:expired"`
}

func (User) TableName() string {
	return "user"
}

func (UserBillInfo) TableName() string {
	return "user_bill_info"
}
