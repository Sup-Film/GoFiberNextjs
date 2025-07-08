package main

import (
	"fiber-restapi/controllers"
	"fiber-restapi/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// สร้าง service และ controller สำหรับผู้ใช้
	userService := services.NewUserService()                     // สร้าง UserService ใหม่
	userController := controllers.NewUserController(userService) // สร้าง UserController ใหม่

	// กำหนด Route สำหรับการลงทะเบียนผู้ใช้
	app.Post("/register", userController.RegisterHandlers) // ลงทะเบียนผู้ใช้ใหม่

	// กำหนด Route สำหรับการเข้าสู่ระบบผู้ใช้
	app.Post("/login", userController.LoginHandler) // เข้าสู่ระบบผู้ใช้

	// Start server
	app.Listen(":3000")
}
