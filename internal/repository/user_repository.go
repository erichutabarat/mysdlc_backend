package repository

import (
	"mysdlc_backend/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct { DB *gorm.DB }

func (r *UserRepository) CreateUser(u *model.User) error {
    return r.DB.Create(u).Error
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	user.Password = "" // Clear password before returning
	return &user, nil
}