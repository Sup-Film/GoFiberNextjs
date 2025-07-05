package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// สร้างแอปพลิเคชัน Fiber
	app := fiber.New()

	// * Basic Routing
	// กำหนด Route สำหรับหน้าแรก
	app.Get("/", func(c *fiber.Ctx) error {
		// ส่งข้อความ "Hello, World!" กลับไปยังผู้ใช้
		return c.SendString("Hello, World!")
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		// ส่งข้อความ "About Page" กลับไปยังผู้ใช้
		return c.SendString("About Page")
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		// ส่งข้อความ "About Page" กลับไปยังผู้ใช้
		return c.SendString("Test Page")
	})

	// * Dynamic Routing
	// กำหนด Route ที่มีพารามิเตอร์
	app.Get("/user/:id", func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "id" จาก URL
		id := c.Params("id")
		return c.SendString("User ID: " + id)
	})

	// กำหนด Route ที่มีพารามิเตอร์หลายตัว
	app.Get("/user/:id/:name", func(c *fiber.Ctx) error {
		// ดึงค่าพรารามิเตอร์ "id" และ "name" จาก URL
		id := c.Params("id")
		name := c.Params("name")
		return c.SendString("User ID: " + id + ", Name: " + name)
	})

	// กำหนด Route ที่มีพารามิเตอร์แบบ Optional
	// ตัวอย่าง URL: http://localhost:3000/optionalparam/
	// หรือ http://localhost:3000/optionalparam/123
	app.Get("/optionalparam/:id?", func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "id" จาก URL ถ้ามี
		id := c.Params("id")
		return c.SendString("Optional ID: " + id)
	})

	// * Query Parameters String
	// กำหนด Route ที่รับพารามิเตอร์แบบ Query String
	// ตัวอย่าง URL: http://localhost:3000/search?q=golang
	app.Get("/search", func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "q" จาก Query String
		query := c.Query("q")
		return c.SendString("Search Query: " + query)
	})

	// กำหนด Route ที่รับพารามิเตอร์แบบ Query String หลายตัว
	// ตัวอย่าง URL: http://localhost:3000/search?name=John&age
	app.Get("/search/multiple", func(c *fiber.Ctx) error {
		query := c.Query("q")
		page := c.Query("page")
		return c.SendString("Search Query: " + query + ", Page: " + page)
	})

	// * กำหนด Route แบบ Wildcard
	// กำหนด Route ที่รับพารามิเตอร์แบบ Wildcard
	// ตัวอย่าง URL: http://localhost:3000/files/images/photo.jpg/abc/def/ghi
	app.Get("/wildcard/*", func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ Wildcard จาก URL
		wildcard := c.Params("*")
		return c.SendString("Wildcard Path: " + wildcard)
	})

	// * การกำหนด Constraints เพื่อบอกว่า Route นี้ต้องการพารามิเตอร์ที่เป็นประเภทใดและมีรูปแบบอย่างไร
	// กำหนด Route ที่มีพารามิเตอร์ที่ต้องเป็นตัวเลข
	// ตัวอย่าง URL: http://localhost:3000/product/123 (ตัวเลขเท่านั้น)
	// วิธีกำหนดให้ใช้ <type> เช่น <int> หรือ <string> ต่อท้ายพารามิเตอร์
	// และสามารถใช้ Regular Expression เพื่อกำหนดรูปแบบได้ เช่น <regex:[a-z]+>
	// หรือใช้ <min(5)> เพื่อกำหนดค่าต่ำสุด
	app.Get("/constraints/:id<int>", func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "id" จาก URL
		id := c.Params("id")
		return c.SendString("Product ID: " + id)
	})

	// เริ่มต้นเซิร์ฟเวอร์ที่พอร์ต 3000
	app.Listen(":3000")
}
