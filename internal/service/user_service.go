package service

import (
    "mysdlc_backend/internal/model"
    "mysdlc_backend/internal/repository"
    "golang.org/x/crypto/bcrypt"
)

type UserService struct { Repo *repository.UserRepository }

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