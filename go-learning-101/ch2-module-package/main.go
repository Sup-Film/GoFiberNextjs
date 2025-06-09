package main

import (
	"fmt"
	// วิธีการ Import package แต่ยังไม่ใช้งาน
	// ใช้ _ เพื่อบอกว่าเราจะไม่ใช้ package นี้ในตอนนี้แต่จะใช้ในอนาคต
	_ "math"

	"github.com/Sup-Film/samplego/greeting"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("CH2 - Module and Package")

	// Generate a new UUID
	// สร้างตัวแปรเพื่อเก็บ UUID
	var id string = uuid.New().String()
	fmt.Println("Generated UUID:", id)

	// เรียกใช้ฟังก์ชันจาก package greeting
	greeting.SayGreeting()
	// เรียกใช้ฟังก์ชันจาก package greeting ที่เป็น private
	greeting.SayMeeting()
}
