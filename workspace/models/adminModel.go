package models

import (
	"utils"
)

// 数据库实例

type Admin struct {
	Id        int             `gorm:"primary_key" column:"id"`
	Username  string          `gorm:"not null;unique" column:"username"`
	Password  string          `gorm:"not null" column:"password"`
	Role      int             `gorm:"not null" column:"role"`
	CreatedAt utils.LocalTime `gorm:"type:datetime not null" column:"created_at" `
	UpdatedAt utils.LocalTime `gorm:"type:datetime not null" column:"updated_at"`
	AdminRole AdminRole       `gorm:"foreignKey:Role;"`
}

type AdminRole struct {
	Id   int
	Name string `gorm:"not null;unique;column:name"`
}

type AdminUserToken struct {
	Id        int             `gorm:"primary_key"`
	UserId    int             `gorm:"not null;column:userid"`
	Token     string          `gorm:"not null;column:token"`
	CreateAt  utils.LocalTime `gorm:"type:datetime not null;column:created_at"`
	ExpiredAt utils.LocalTime `gorm:"type:datetime not null;column:expired_at"`
	Expired   int             `gorm:"not null;default:0"`
}

func (Admin) TableName() string {
	return "admin_user"
}

func (AdminRole) TableName() string {
	return "admin_role"
}

func (AdminUserToken) TableName() string {
	return "admin_user_token"
}
