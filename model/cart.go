package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserId    uint `gorm:"not null"`
	ProductId uint `gorm:"not null"`
	BossId    uint `gorm:"not null"`
	Num       int  `gorm:"not null"`
	MaxNum    int  `gorm:"not null"`
	Check     bool
}
