package model

import "gorm.io/gorm"

type User struct {
	Model
	Name string `gorm:"size:100;not null;default:''"`
	Email string `gorm:"size:200;not null;default:'';uniqueIndex"`
	Phone string `gorm:"size:20;not null;default:'';uniqueIndex"`
	Password string `gorm:"size:200;not null;default:''"`
	DeletedAt gorm.DeletedAt `gorm:"type:int(11)"`
}