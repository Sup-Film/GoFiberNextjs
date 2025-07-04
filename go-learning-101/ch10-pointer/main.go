package main

import "fmt"

func main() {
	SamplePointer()
	PointerSlice()
	fmt.Println("-----------------------------")
	// * Function ที่รับ Pointer เป็น Argument
	age := 20
	// การส่งค่าเข้าไปในฟังก์ชั่นที่รับพารามิเตอร์เป็น Pointer เราต้องส่งเป็นที่อยู่ของตัวแปรนั้น ๆ เข้าไป
	ChangeAge(&age)

	fmt.Println("------------------------------")

	// * กรณีใช้ Pointer กับ Struct
	// เรียกใช้ฟังก์ชัน SamplePointer2 เพื่อเปลี่ยนสีเตียงของห้อง
	// โดยส่ง Pointer ของ Struct Room เข้าไป
	SamplePointer2()

}
