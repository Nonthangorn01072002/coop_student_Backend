package services

import (
	"coop_student_backend/internal/domain"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserLoginService struct {
	db *gorm.DB
}

func NewUserLoginService(database *gorm.DB) *UserLoginService {
	return &UserLoginService{db: database}
}

func (s *UserLoginService) Create(Username string, Password string) (*domain.UserLogin, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	userLogin := &domain.UserLogin{
		Username: Username,
		Password: string(hashPassword),
	}
	if err := s.db.Create(userLogin).Error; err != nil {
		return nil, err
	}
	return userLogin, nil
}
