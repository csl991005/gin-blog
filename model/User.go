package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20);not null" json:"user_name"`
	PassWord string `gorm:"type:varchar(500);not null" json:"pass_word"`
	Role     int    `gorm:"type:int" json:"role"`
}
