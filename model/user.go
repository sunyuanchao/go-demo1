package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"varchar(20);not null"`
	Password  string `gorm:"size:255;not null"`
	Telephone string `gorm:"varchar(200);not null;unique"`
}
