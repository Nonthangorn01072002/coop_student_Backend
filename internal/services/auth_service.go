package services

import (
	"coop_student_backend/internal/domain"
	"coop_student_backend/internal/dto"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}


func (a *AuthService) ExtractUserIDFromJWT(tokenStr string) (int, error) {
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecretKey), nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims type")
	}

	userIdFloat, ok := claims["userId"].(float64)
	if !ok {
		return 0, errors.New("userId not found")
	}

	return int(userIdFloat), nil
}

func NewAuthService(database *gorm.DB) *AuthService {
	return &AuthService{db: database}
}

func (a *AuthService) Login(loginDto dto.LoginDto) (string, error) {
	var existingUserlogin domain.UserLogin
	if err := a.db.First(&existingUserlogin, "username = ?", loginDto.Username).Error; err != nil {
		return "", err
	}

	comparePasswordErr := bcrypt.CompareHashAndPassword([]byte(existingUserlogin.Password), []byte(loginDto.Password))
	if comparePasswordErr != nil {
		return "", comparePasswordErr
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	tokenSigner := jwt.New(jwt.SigningMethodHS256)
	claims := tokenSigner.Claims.(jwt.MapClaims)
	claims["userId"] = existingUserlogin.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token, signTokenErr := tokenSigner.SignedString([]byte(jwtSecretKey))
	if signTokenErr != nil {
		return "", signTokenErr
	}
	return token, nil
}

