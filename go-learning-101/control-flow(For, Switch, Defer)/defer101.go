package main

import (
	"fmt"
)

func Defer101() {
	defer fmt.Println("คำสั่งนี้จะถูกเรียกใช้เมื่อฟังก์ชันนี้สิ้นสุด")
	fmt.Println("คำสั่งนี้จะถูกเรียกใช้ก่อนคำสั่ง defer")
}

func StackingDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
