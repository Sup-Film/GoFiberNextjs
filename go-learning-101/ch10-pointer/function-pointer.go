package main

import (
	"fmt"
)

// Function ที่รับ Pointer เป็น Argument
func ChangeAge(age *int) {
	fmt.Println("แสดงที่อยู่ของ Age:", age)
	fmt.Println("แสดงค่า Age ที่รับเข้ามาจากอากิวเมนต์:", *age)

	*age += 1 // เพิ่มค่า Age ขึ้น 1
	fmt.Println("แสดงค่า Age หลังจากเพิ่มขึ้น 1:", *age)
}

// * กำหนดโครงสร้าง Room
// * เป็นการกำหนดโครงสร้างข้อมูลใหม่ ที่ใช้ชื่อชนิดข้อมูลว่า Room
// * โดยมีฟิลด์ RoomNumber และ BedColor
type Room struct {
	RoomNumber int
	BedColor   string
}

func ChangeBedColorWithPointer(room *Room, newColor string) {
	fmt.Println("เปลี่ยนสีเตียงของห้องที่ชี้ไปที่หมายเลข:", room.RoomNumber)

	// เปลี่ยนสีเตียงของห้องที่ชี้ไป
	// ในกรณีที่เป็น Struct เราจะไม่ต้องใช้ * เพื่อเข้าถึงฟิลด์ภายใน Struct
	room.BedColor = newColor
}

// * กรณีใช้ Pointer
func SamplePointer2() {
	// สร้าง Room หมายเลข 101 และสีเตียง "Blue"
	room101 := Room{
		RoomNumber: 101,
		BedColor:   "Blue",
	}

	fmt.Println("ห้องหมายเลข:", room101.RoomNumber)
	fmt.Println("สีเตียงเดิม:", room101.BedColor)

	// เรียกใช้ฟังก์ชัน ChangeBedColorWithPointer โดยส่ง Pointer ของ room101
	ChangeBedColorWithPointer(&room101, "Red")
	fmt.Println("สีเตียงหลังจากเปลี่ยน:", room101.BedColor)
}
