package controllers

import (
	"fiber-restapi/models"
	"fiber-restapi/services" // แก้ไขชื่อแพ็กเกจให้ตรงกับที่สร้างไว้

	"github.com/gofiber/fiber/v2"
)

// สร้าง UserController struct สำหรับเก็บ UserService
type UserController struct {
	UserService *services.UserService
}

// สร้าง NewUserController ฟังก์ชั่นสำหรับสร้าง UserController ใหม่
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// RegisterHandlers สำหรับการลงทะเบียนผู้ใช้ใหม่
func (ctrl *UserController) RegisterHandlers(c *fiber.Ctx) error {
	user := new(models.User) // สร้างตัวแปร user ใหม่จากโมเดล User
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "ข้อมูลไม่ถูกต้อง",
			"error":   "Invalid request body",
		})
	}

	registeredUser, err := ctrl.UserService.RegisterUser(*user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "ลงทะเบียนไม่สำเร็จ",
			"error":   err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "ลงทะเบียนสำเร็จ",
		"user": fiber.Map{
			"id":         registeredUser.ID,
			"username":   registeredUser.Username,
			"email":      registeredUser.Email,
			"fullname":   registeredUser.Fullname,
			"role":       registeredUser.Role,
			"created_at": registeredUser.CreatedAt,
			"updated_at": registeredUser.UpdatedAt,
		},
	})
}

// LoginHandler สำหรับการเข้าสู่ระบบผู้ใช้
func (ctrl *UserController) LoginHandler(c *fiber.Ctx) error {
	user := new(models.User) // สร้างตัวแปร user ใหม่จากโมเดล User
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "ข้อมูลไม่ถูกต้อง",
			"error":   "Invalid request body",
		})
	}

	loggedInUser, token, err := ctrl.UserService.Login(user.Username, user.Password)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "เข้าสู่ระบบไม่สำเร็จ",
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "เข้าสู่ระบบสำเร็จ",
		"user": fiber.Map{
			"id":         loggedInUser.ID,
			"username":   loggedInUser.Username,
			"email":      loggedInUser.Email,
			"fullname":   loggedInUser.Fullname,
			"role":       loggedInUser.Role,
			"created_at": loggedInUser.CreatedAt,
			"updated_at": loggedInUser.UpdatedAt,
		},
		"token": token,
	})
}
