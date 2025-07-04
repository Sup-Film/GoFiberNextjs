package main

import "fmt"

func SampleMap() {
	// การประกาศตัวแปรแบบ map จะแบ่งเป็น key-value
	// var ตัวแปร map[ประเภทของ key]ประเภทของ value

	m0 := make(map[string]int)                  // ประกาศตัวแปร m0 เป็น map ที่มี key เป็น string และ value เป็น int
	m0["apple"] = 5                             // กำหนดค่า key "apple" ให้มี value เป็น 5
	m0["banana"] = 3                            // กำหนดค่า key "banana" ให้มี value เป็น 3
	m0["orange"] = 2                            // กำหนดค่า key "orange" ให้มี value เป็น 2
	fmt.Println("Map m0:", m0)                  // แสดงผล map m0
	fmt.Println("Map m0[apple]: ", m0["apple"]) // แสดงผล value ของ key "apple" ใน map m0

	// *คำสั่งตรวจสอบว่ามี key อยู่ใน map หรือไม่
	if value, exists := m0["banana"]; exists {
		fmt.Println("Key 'banana' exists with value:", value) // แสดงผล value ของ key "banana" ถ้ามีอยู่
	} else {
		fmt.Println("Key 'banana' does not exist") // แสดงข้อความถ้า key "banana" ไม่มีอยู่ใน map
	}

	// * ลบสมาชิกใน map
	delete(m0, "apple")                              // ลบ key "apple" ออกจาก map m0
	fmt.Println("Map m0 หลังจากลบ Key 'apple':", m0) // แสดงผล map m0 หลังจากลบ key "apple"

	// * วนลูปข้อมูลใน map
	for key, value := range m0 {
		fmt.Println("Key:", key, "Value:", value) // แสดงผล key และ value ของแต่ละสมาชิกใน map m0
	}

	// * การสร้าง map โดยใช้ literal
	grade := map[string]int{
		"John":  90,
		"Jane":  85,
		"Bob":   78,
		"Alice": 92,
	}

	fmt.Println("Map grade:", grade) // แสดงผล map grade
}
