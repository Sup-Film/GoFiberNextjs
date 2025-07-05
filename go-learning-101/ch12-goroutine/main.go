package main

import (
	"fmt"
	"time"
)

func brewCoffee() {
	fmt.Println("Starting to brew coffee...")
	time.Sleep(2 * time.Second) // จำลองการทำงานที่ใช้เวลา 2 วินาที
	fmt.Println("Coffee is ready!")
}

func bakeBread() {
	fmt.Println("Starting to bake bread...")
	time.Sleep(3 * time.Second) // จำลองการทำงานที่ใช้เวลา 3 วินาที
	fmt.Println("Bread is ready!")
}

func washDishes() {
	fmt.Println("Starting to wash dishes...")
	time.Sleep(1 * time.Second) // จำลองการทำงานที่ใช้เวลา 1 วินาที
	fmt.Println("Dishes are clean!")
}

func main() {
	fmt.Println("Goroutine Example")
	fmt.Println("--------------------")
	go brewCoffee() // เริ่มทำกาแฟใน goroutine
	go bakeBread()  // เริ่มทำขนมปังใน goroutine
	go washDishes() // เริ่มล้างจานใน goroutine

	time.Sleep(5 * time.Second) // รอให้ goroutine ทำงานเสร็จ (ในกรณีนี้ 5 วินาทีเพียงพอสำหรับการทำงานทั้งหมด)
	fmt.Println("All tasks are done!")

	// หมายเหตุ: ในการใช้งานจริง ควรใช้ sync.WaitGroup แทนการใช้ time.Sleep เพื่อรอให้ goroutine ทั้งหมดทำงานเสร็จ
}
