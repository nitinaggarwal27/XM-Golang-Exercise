package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null;unique;" binding:"required"`
	Code    string `json:"code" gorm:"not null;unique;" binding:"required"`
	Country string `json:"country" binding:"required"`
	Website string `json:"website"`
	Phone   string `json:"phone"`
}
