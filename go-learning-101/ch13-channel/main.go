package main

import (
	"fmt"
	"time"
)

func main() {
	// * สร้าง channel ในการสื่อสารระหว่าง goroutine
	ch := make(chan string)

	// * Annonymous function เป็นการประกาศฟังก์ชันที่ไม่มีชื่อ และจะถูกเรียกใช้ทันที
	// เรียกใช้ Goroutine เพื่อส่งข้อมูลไปยัง channel
	go func() {
		time.Sleep(1 * time.Second)   // จำลองการทำงานที่ใช้เวลา 2 วินาที
		ch <- "Hello From Goroutine!" // ส่งข้อความไปยัง channel
	}()

	msg := <-ch      // รอรับข้อความจาก channel
	fmt.Println(msg) // แสดงข้อความที่ได้รับ
}
