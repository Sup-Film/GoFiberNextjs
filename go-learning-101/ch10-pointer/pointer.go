package main

import "fmt"

// * Pointer เป็นตัวแปรที่เก็บที่อยู่ของตัวแปรอื่น ๆ
func SamplePointer() {
	value := 42 // ประกาศตัวแปร value ที่เก็บค่า 42

	pointer := &value // สร้่าง pointer ที่เก็บที่อยู่ของ value

	// * แสดงค่าของ value และ pointer
	fmt.Println("Value:", value)               // แสดงค่า 42
	fmt.Println("Pointer:", pointer)           // แสดงตำแหน่ง address ของ value เช่น0xc000012088
	fmt.Println("Value in Pointer:", *pointer) // แสดง value ที่ pointer เก็บ address อยู่

	fmt.Println("-----------------------------")

	// * การเปลี่ยนแปลงค่าผ่าน pointer
	// ต้องใช้ *pointer เพื่อเข้าถึงค่าในที่อยู่ที่ pointer ชี้ไป
	*pointer = 100 // เปลี่ยนค่า value ผ่าน pointer

	// * แสดงค่าหลังจากเปลี่ยนแปลง
	fmt.Println("New Value:", value)                        // แสดงค่า 100
	fmt.Println("Pointer after change:", pointer)           // แสดงตำแหน่ง address ของ value เช่น0xc000012088
	fmt.Println("Value in Pointer after change:", *pointer) // แสดง value ที่ pointer เก็บ address อยู่
}
