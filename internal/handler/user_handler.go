package handler

import (
	"coop_student_backend/internal/domain"
	"coop_student_backend/internal/dto"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service domain.UserService
}

func NewUserHandler(s domain.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUserById(c *fiber.Ctx) error{
	id := c.Params("id")
	user, err:= h.service.FindById(id)
	if err != nil{
		return c.Status(500).JSON(fiber.Map{
			"message":"Failed to get user by ID.",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message":"Get user by ID sucessfuly",
		"result":user,
	})
}

func (h *UserHandler) GetAllUser(c *fiber.Ctx) error {
	uidStr := c.Cookies("coopToken")
	if uidStr == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized, no Token found",
		})
	}

	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid coopToken format",
		})
	}

	users, err := h.service.FindAll(uid)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to get user.",
		})
	}

	if users == nil {
		return c.Status(403).JSON(fiber.Map{
			"message": " You don't have permission to view",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Get all user successfully",
		"result":  users,
	})
}


func (h *UserHandler) CreateUser(c *fiber.Ctx) error{
	var createUserDto dto.CreateUserDto
	if err := c.BodyParser(&createUserDto); err != nil{
		return c.Status(400).JSON(fiber.Map{
			"message":"Request body is invalid",
		})
	}
	user, err := h.service.Create(createUserDto)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message":"Fail to create user",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message":"Create user sucessfuly",
		"result":user,
	})
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error{
	id := c.Params("id")
	var updateUser dto.UpdateUserDto
	if err := c.BodyParser(&updateUser); err != nil{
		return c.Status(400).JSON(fiber.Map{
			"message":"Request body is invalid",
		})
	}
	upProduct, err := h.service.Update(id,updateUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message":"Fail to create user",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message":"Create user by ID sucessfuly",
		"result":upProduct,
	})
}

func (h *UserHandler) DeleteUserById(c *fiber.Ctx) error{
	id := c.Params("id")
	user, err:= h.service.Delete(id)
	if err != nil{
		return c.Status(500).JSON(fiber.Map{
			"message":"Failed to delete user by ID.",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message":"Delete user by ID sucessfuly",
		"result":user,
	})
}