package main

import (
	"strconv"

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

	// * การกำหนด Route แบบ Regex
	// กำหนด Route ที่มีพารามิเตอร์ที่ต้องเป็นตัวเลข 5 หลัก
	// ตัวอย่าง URL: http://localhost:3000/regex/12345 (ตัวเลข 5 หลักเท่านั้น)
	// ใช้ Regular Expression เพื่อกำหนดรูปแบบพารามิเตอร์
	// รูปแบบ: :parameter<regex(pattern)>
	app.Get(`/regex/:id<regex(\d{5})>`, func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "id" จาก URL
		id := c.Params("id")
		return c.SendString("Product ID: " + id)
	})

	// กำหนด Route เป็นรูปแบบวันที่ YYYY-MM-DD (ต้องตรงทุกตัวอักษร)
	// ตัวอย่าง URL: http://localhost:3000/date/2025-02-02 (ต้องตรง 10 ตัวอักษรพอดี)
	// ใช้ ^ และ $ เพื่อบังคับให้ match ทั้งหมด
	// รูปแบบ: :parameter<regex(^pattern$)>
	app.Get(`/date/:date<regex(^\d{4}-\d{2}-\d{2}$)>`, func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "date" จาก URL
		date := c.Params("date")
		return c.SendString("Date: " + date)
	})

	// --------------------------------------------------
	// * การกำหนด Route แบบ HTTP Method ต่าง ๆ
	// รับค่า GET, POST, PUT, DELETE จาก Request

	// กำหนด Struct Person สำหรับการรับข้อมูล JSON
	type Person struct {
		ID    int    `json:"id"` // "id" จะไปแมพกับ JSON key "id" ที่ส่งเข้ามา
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	// ตัวแปรเก็บข้อมูลผู้ใช้ โดยใช้ slice ของ struct Person
	var people []Person = []Person{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
	}

	// GET: ดึงข้อมูลผู้ใช้ทั้งหมด
	// ตัวอย่าง URL: http://localhost:3000/users
	app.Get("/users", func(c *fiber.Ctx) error {
		// ส่งข้อมูลผู้ใช้ทั้งหมดกลับไปยังผู้ใช้ในรูปแบบ JSON
		return c.JSON(fiber.Map{
			"message": "All users retrieved successfully",
			"data":    people,
			"count":   len(people),
		})
	})

	// POST: สร้างผู้ใช้ใหม่
	// ตัวอย่าง URL: http://localhost:3000/users
	// ตัวอย่าง JSON ที่ส่งมาใน Request Body:
	// {
	//   "id": 2,
	//   "name": "Jane Doe",
	//   "email": "jane@example.com"
	// }
	// ใช้ c.BodyParser() เพื่อแปลง JSON ที่ส่งมาเป็น struct Person
	app.Post("/users", func(c *fiber.Ctx) error {
		// สร้างตัวแปรสำหรับเก็บผู้ใช้ใหม่ เป็นตัวแปร Pointer ที่ชี้ไปยัง struct Person
		person := new(Person)
		// แปลง JSON ที่ส่งมาใน Body เป็น struct Person
		// ถ้าเกิดข้อผิดพลาดในการแปลง JSON ให้ส่งข้อความผิดพลาดกลับไป
		if err := c.BodyParser(person); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request data",
				"error":   err.Error(),
			})
		}

		// เพิ่มผู้ใช้ใหม่เข้าไปใน slice ของผู้ใช้
		// ใช้ Pointer เพื่อนำค่าใน struct เข้าไปใน slice
		people = append(people, *person)

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "User created successfully",
			"data":    person,
			"count":   len(people),
		})
	})

	// GET: ดึงข้อมูลผู้ใช้ตาม ID
	// ตัวอย่าง URL: http://localhost:3000/users/1
	// ถ้าไม่พบผู้ใช้ที่มี ID ตรงกับที่ระบุ ให้ส่งข้อความผิดพลาดกลับไป
	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		// แปลง ID เป็นตัวเลข
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid user ID",
			})
		}

		// ค้นหาผู้ใช้ที่มี ID ตรงกับที่ระบุ เนื่องจากเราใช้เป็น slice เราต้องวนลูปค่าเพื่อหา ID
		for _, person := range people {
			if person.ID == idInt {
				return c.Status(fiber.StatusFound).JSON(fiber.Map{
					"message": "User retrieved successfully",
					"data":    person,
				})
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"data":    nil,
		})
	})

	// UPDATE: อัปเดตข้อมูลผู้ใช้ตาม ID
	// ตัวอย่าง URL: http://localhost:3000/users/1
	// ถ้าไม่พบผู้ใช้ที่มี ID ตรงกับที่ระบุ ให้ส่งข้อความผิดพลาดกลับไป
	// ต้องส่งข้อมูล JSON ที่ต้องการอัปเดตใน Request Body
	// {
	//   "name": "John Smith",
	//   "email": "john.smith@example.com"
	// }
	app.Put("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		// แปลง ID เป็นตัวเลข
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid user ID",
			})
		}

		// สร้างตัวแปรสำหรับเก็บข้อมูลผู้ใช้ใหม่
		person := new(Person)

		// แปลง JSON ที่ส่งมาใน Body เป็น struct Person
		if err := c.BodyParser(person); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request data",
			})
		}

		// หาผู้ใช้ที่มี ID ตรงกับที่ระบุ
		for i, p := range people {
			if p.ID == idInt {
				// อัปเดตข้อมูล
				// นำค่าของ person ที่เก็บค่าจาก JSON Body มาใช้
				// people[i] เป็นค่าที่ชี้ไปยัง slice ของผู้ใช้ที่ตรงกับ ID ที่ส่งมาใน Request params
				people[i].Name = person.Name
				people[i].Email = person.Email
				return c.Status(fiber.StatusOK).JSON(fiber.Map{
					"message": "User updated successfully",
					"data":    people[i], // ส่งข้อมูลผู้ใช้ที่อัปเดตกลับไป
				})
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	})

	// DELETE: ลบข้อมูลผู้ใช้ตาม ID
	// ตัวอย่าง URL: http://localhost:3000/users/1
	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		// แปลง ID เป็นตัวเลข
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid user ID",
			})
		}

		// หาผู้ใช้ที่มี ID ตรงกับที่ระบุ
		for i, person := range people {
			if person.ID == idInt {
				// ลบผู้ใช้ออกจาก slice ของผู้ใช้
				// Go จะไม่มี method สำหรับลบค่าออกจาก slice โดยตรงเหมือนกับ Array ในภาษาอื่น ๆ
				people = append(people[:i], people[i+1:]...)
				return c.Status(fiber.StatusOK).JSON(fiber.Map{
					"message": "User deleted successfully",
					"data":    person,
				})
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	})

	// เริ่มต้นเซิร์ฟเวอร์ที่พอร์ต 3000
	app.Listen(":3000")
}
