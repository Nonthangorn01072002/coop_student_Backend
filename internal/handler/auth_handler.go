package handler

import (
	"coop_student_backend/internal/domain"
	"coop_student_backend/internal/dto"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service domain.AuthService
}

func NewAuthHandler (s domain.AuthService) *AuthHandler{
	return &AuthHandler{service : s}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error{
	var loginDto dto.LoginDto
	if err := c.BodyParser(&loginDto); err != nil{
		return c.Status(400).JSON(fiber.Map{"error":"Invalid request"})
	}

	token, err := h.service.Login(loginDto)
	if err!= nil{
		return c.Status(500).JSON(fiber.Map{"error" : "Fail to login"})
	}

	c.Cookie(&fiber.Cookie{
		Name:"coopToken",
		Value: token,
		Expires: time.Now().Add(time.Hour *72),
		SameSite :"Lax",
		Secure : false,
		HTTPOnly: true,
	})

	return c.Status(200).SendString("Login Successfully")
}