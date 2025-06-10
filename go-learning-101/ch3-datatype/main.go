package main

import (
	"fmt"
	"html" // แปลง String ให้เป็น HTML
)

func main() {
	fmt.Println("Data Types in Go")
	// มี 2 ประเภทหลัก ๆ คือ
	// 1. Primitive Types (ประเภทพื้นฐาน)
	// 2. Composite Types (ประเภทประกอบ)
	// Primitive Types ได้แก่ (int, float, bool, string)
	// Composite Types ได้แก่ (array, slice, map, struct, interface)

	// Primitive Types
	var a int = 10 // int8, int16, int32, int64
	fmt.Println("INT:", a)

	var b uint = 20 // uint8, uint16, uint32, uint64
	fmt.Println("UINT:", b)

	var c float64 = 3.14 // float32, float64
	fmt.Println("FLOAT:", c)

	var d bool = true // true, false
	fmt.Println("BOOL:", d)

	// String Data Type
	// Interpreted strings เป็น string ที่มีการแปลความหมายเป็นตัวอักษร
	var e string = "Hello, Go! \nThis is Interpreted String"
	fmt.Println("STRING:", e)

	// Raw strings เป็น string ที่ไม่ต้องการการแปลความหมาย เช่น การใช้ backticks
	// Raw strings ถ้าเราใส่อะไรเข้าไปภายในจะไม่มีการแปลความหมายเช่นถ้าใส่ \n ก็จะไม่ขึ้นบรรทัดใหม่
	var f string = `Hello\nGo!`
	fmt.Println("RAW STRING:", f)

	// Escape sequences (อักขระพิเศษ)
	// \n = New Line (ขึ้นบรรทัดใหม่)
	// \t = Tab (เว้นวรรค)
	// \\ = Backslash (เครื่องหมายแบ็คสแลช)
	// \" = Double Quote (เครื่องหมายคำพูดคู่)
	// \' = Single Quote (เครื่องหมายคำพูดเดี่ยว)
	// \r = Carriage Return (ย้ายเคอร์เซอร์ไปยังจุดเริ่มต้นของบรรทัด)
	// \v = Vertical Tab (เว้นวรรคแนวตั้ง)
	// \b = Backspace (ลบตัวอักษรก่อนหน้า)

	// วิธีการเปลี่ยน Raw String เป็น Interpreted String
	// ต้องใช้ package html
	var tag string = `<a href="https://example.com">Click Here</a>`
	// แปลงเป็น HTML ที่ปลอดภัยเพื่อป้องกัน XSS (Cross-Site Scripting)
	var str_esc = html.EscapeString(tag)         // แปลงเป็น HTML ที่ปลอดภัย
	var str_unesc = html.UnescapeString(str_esc) // แปลง HTML ที่ปลอดภัยกลับเป็นข้อความปกติ
	fmt.Println("EscapeString:", str_esc)
	fmt.Println("Unescaped HTML TAG:", str_unesc)
}
