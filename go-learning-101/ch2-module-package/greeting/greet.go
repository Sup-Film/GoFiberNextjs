package greeting

import (
	"fmt"

	personal "github.com/Sup-Film/samplego/greeting/internal"
)

func SayGreeting() {
	fmt.Println("Hello, From Greeting Package")

	// เรียกใช้ฟังก์ชันที่เป็น private ของ package นี้
	sayhi()

	// เรียกใช้ฟังก์ชันจาก internal
	personal.SayPersonal()
}
