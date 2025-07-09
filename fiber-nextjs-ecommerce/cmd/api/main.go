package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create fiber instance
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, From Fiber! Next Ecommerce API",
		})
	})

	// Start the server
	log.Fatal(app.Listen(":3000"))
	log.Println("Server is running on http://localhost:3000")
}
