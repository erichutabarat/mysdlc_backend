package repository

import (
	"mysdlc_backend/internal/model"
	"gorm.io/gorm"
	"errors"
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

// Project Member
func (r *UserRepository) GetInvitationsByUserID(userID uint) ([]model.Invitation, error) {
	var invitations []model.Invitation
	err := r.DB.Table("project_members").
		Select("project_members.project_id, project_members.user_id, project_members.role, project_members.status").
		Where("project_members.user_id = ?", userID).
		Scan(&invitations).Error
	return invitations, err
}

func (r *UserRepository) UpdateInvitationStatus(userID uint, projectID uint, status model.MemberStatus) error {
    result := r.DB.Table("project_members").
        Where("user_id = ? AND project_id = ? AND status = ?", userID, projectID, "pending").
        Update("status", status)

    if result.Error != nil {
        return result.Error
    }

    if result.RowsAffected == 0 {
        return errors.New("no pending invitation found for this user and project")
    }

    return nil
}