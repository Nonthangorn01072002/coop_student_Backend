package main

import (
	"log"
	"os"

	"coop_student_backend/internal/config"
	"coop_student_backend/internal/handler"
	"coop_student_backend/internal/middlewares"
	"coop_student_backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Env not found")
	}

	app := fiber.New()

	db := config.InitDb()

	userLoginService := services.NewUserLoginService(db)
	userService := services.NewUserService(db,userLoginService)
	userHandler := handler.NewUserHandler(userService)
	userGroup := app.Group("/api/user")

	userGroup.Use(middlewares.ValidateJwtMiddleware)

	userGroup.Get("/",userHandler.GetAllUser)
	userGroup.Get("/:id",userHandler.GetUserById)
	app.Post("/api/user",userHandler.CreateUser)
	userGroup.Post("/:id",userHandler.UpdateUser)
	userGroup.Delete("/:id",userHandler.DeleteUserById)

	authService := services.NewAuthService(db)
	authHandler := handler.NewAuthHandler(authService)
	app.Post("/api/auth/login", authHandler.Login)


	app.Get("/api/health-check", func (c *fiber.Ctx) error {
        return c.Status(200).SendString("Health check is ok!")
    })

	port :=os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}
	app.Listen(":"+port)
}