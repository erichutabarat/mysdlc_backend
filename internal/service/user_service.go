package service

import (
    "errors"
    "mysdlc_backend/internal/model"
    "mysdlc_backend/internal/repository"
    "mysdlc_backend/pkg/jwt"
    "golang.org/x/crypto/bcrypt"
)

type UserService struct { Repo *repository.UserRepository }

type UserServiceInterface interface {
	Register(u *model.User) error
	Login(email, password string) (string, error)
	Profile(userID uint) (*model.User, error)
	Invitation(userID uint) ([]model.Invitation, error)
	RespondInvitation(userID uint, req model.RespondInvitationRequest) error
}

// safety check to ensure UserService implements UserServiceInterface
var _ UserServiceInterface = (*UserService)(nil)

func (s *UserService) Register(u *model.User) error {
    // 1. Business Logic: Set Default Role
    if u.Role == "" {
        u.Role = "user"
    }

    // 2. Business Logic: Hash Password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(hashedPassword)

    // 3. Persist
    return s.Repo.CreateUser(u)
}

func (s *UserService) Login(email, password string) (string, error) {
    // 1. Fetch User
    user, err := s.Repo.GetUserByEmail(email)
    if err != nil {
        return "", err
    }

    // Verify Password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", errors.New("invalid email or password")
    }

    // Generate JWT
    return jwt.GenerateToken(user.ID, user.Role)
}

func (s *UserService) Profile(userID uint) (*model.User, error) {
    return s.Repo.GetUserByID(userID)
}

func (s *UserService) Invitation(userID uint) ([]model.Invitation, error) {
    return s.Repo.GetInvitationsByUserID(userID)
}

func (s *UserService) RespondInvitation(userID uint, req model.RespondInvitationRequest) error {
    return s.Repo.UpdateInvitationStatus(userID, req.ProjectID, req.Status)
}