package main

import (
	"fmt"
	"log"

	"github.com/Sup-Film/fiber-nextjs-ecommerce/config"
	"github.com/Sup-Film/fiber-nextjs-ecommerce/internal/adapters/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// load configuration
	cfg := config.LoadConfig()

	// Create fiber instance
	app := fiber.New()

	// Connect DB & Auto Migrate Models
	db.Connect(cfg)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, From Fiber! Next Ecommerce API",
		})
	})

	// Start the server
	// การ Listen จะต้องใช้ string format สำหรับพอร์ต เลยต้องใช้ fmt.Sprintf
	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.AppPort)))
	log.Println("Server is running on http://localhost:" + cfg.AppPort)
}
