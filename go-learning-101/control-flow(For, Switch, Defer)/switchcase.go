package main

import (
	"fmt"
)

func SwitchCase() {
	dayOfWeek := 3 // ตัวอย่างการกำหนดวันในสัปดาห์

	switch dayOfWeek {
	case 1:
		fmt.Println("วันนี้คือวันจันทร์")
	case 2:
		fmt.Println("วันนี้คือวันอังคาร")
	case 3:
		fmt.Println("วันนี้คือวันพุธ")
	case 4:
		fmt.Println("วันนี้คือวันพฤหัสบดี")
	case 5:
		fmt.Println("วันนี้คือวันศุกร์")
	case 6:
		fmt.Println("วันนี้คือวันเสาร์")
	case 7:
		fmt.Println("วันนี้คือวันอาทิตย์")
	default:
		fmt.Println("ไม่ใช่วันในสัปดาห์ที่ถูกต้อง")
	}
}

func SwitchCaseCondition() {
	dayOfWeek := "วันจันทร์" // ตัวอย่างการกำหนดวันในสัปดาห์

	switch dayOfWeek {
	case "วันจันทร์", "วันอังคาร":
		fmt.Println("วันนี้เป็นวันทำงาน")
	case "วันพุธ", "วันพฤหัสบดี":
		fmt.Println("วันนี้เป็นวันกลางสัปดาห์")
	case "วันศุกร์":
		fmt.Println("วันนี้เป็นวันศุกร์")
	case "วันเสาร์", "วันอาทิตย์":
		fmt.Println("วันนี้เป็นวันหยุดสุดสัปดาห์")
	default:
		fmt.Println("ไม่ใช่วันในสัปดาห์ที่ถูกต้อง")
	}
}
