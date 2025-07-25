package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateJwtMiddleware(c *fiber.Ctx) error {
	cookie := c.Cookies("coopToken")

	if cookie == ""{
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token)(any,error){
		return []byte(jwtSecretKey),nil
	})

	if err != nil || !token.Valid{
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}