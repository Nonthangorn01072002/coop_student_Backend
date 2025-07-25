package domain

import (
	"coop_student_backend/internal/dto"

	"gorm.io/gorm"
)

type UserLogin struct {
	gorm.Model
	Username string
	Password string
}

type AuthService interface{
	Login(loginDto dto.LoginDto) (string, error)
	ExtractUserIDFromJWT(tokenStr string) (int, error)
}

type UserLoginService interface{
	Create(userLogin UserLogin) (*UserLogin,error)
}

