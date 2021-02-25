package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username; type:varchar(255); not null; unique" json:"username"`
	Password string `gorm:"column:password; type:varchar(255); not null "json:"password"`
	FullName string `gorm:"column:fullname; type:varchar(255); not null "json:"fullname"`
	Photo    string `gorm:"column:photo; type:varchar(255)"json:"photo"`
	RoleID   uint   `gorm:"column:roleId" json:"roleId"`
	Role     Role   `gorm:"foreignKey:RoleID"`
}

type Role struct {
	ID   uint   `gorm:"column:id; primary_key; AUTO_INCREMENT" json:"id"`
	Role string `gorm:"column:role; type:varchar(255); not null" json:"role"`
}
