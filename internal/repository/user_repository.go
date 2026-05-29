package repository

import (
	"mysdlc_backend/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct { DB *gorm.DB }

func (r *UserRepository) CreateUser(u *model.User) error {
    return r.DB.Create(u).Error
}