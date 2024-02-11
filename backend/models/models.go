package models

import "gorm.io/gorm"

type Credentials struct {
	Email    string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
}

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Credentials
}
