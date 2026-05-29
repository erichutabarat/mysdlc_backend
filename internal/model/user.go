package model

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Email    string `gorm:"unique;not null" json:"email"`
    Name     string `gorm:"not null" json:"name"`
    Password string `gorm:"not null" json:"-"`
    Role     string `gorm:"not null;default:'user'" json:"role"`
}