package main

import (
	"fmt"
)

// สร้างฟังก์ชั่นรับ Parameter x และ y โดยให้ฟังก์ชั่นนี้ return ค่าเป็น int
// ในกรณีที่ฟังก์ชั่นรับค่า Parameter ติดกันเป็น Type เดียวกัน สามารถเขียนได้แบบนี้
// func Add(x, y int) int {}
func addNumbers(x int, y int) int {
	fmt.Println("Adding two integers:", x, "+", y)
	sum := x + y
	return sum
}

// Function return multiple values
func swap(x, y int) (int, int) {
	return y, x
}

// ฟังก์ชั่นที่ return ค่าที่ต้องชื่อรอไว้แต่แรก
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // การ return แบบนี้จะใช้ค่าที่กำหนดไว้ในตัวแปร x และ y ที่ประกาศไว้ก่อนหน้า
}

func definition(n1, n2 int) (int, int) {
	sum := n1 + n2
	diff := n1 - n2

	return sum, diff // return ค่า sum และ diff
}

func main() {
	fmt.Println("Result of Add:", addNumbers(10, 20))

	// ในกรณีที่ฟังก์ชั่น return multiple values สามารถรับค่าที่ return ได้หลายค่า
	// หรือจะแสดง return value ทั้งสองค่าเลยก็ได้โดยใช้ fmt.Println(swap(10, 20))
	a, b := swap(10, 20)
	fmt.Println("Swapped values:", a, b)
	fmt.Println(swap(10, 20))

	// เรียกใช้ฟังก์ชั่นที่มีการ return ค่าที่มากกว่า 1 ค่า
	x, y := split(17)
	fmt.Println("Split result:", x, y)
	fmt.Println(split(17))

	// เรียกใช้ฟังก์ชั่นที่ return มากกว่า 1 ค่า
	sum, diff := definition(10, 5)
	fmt.Println("Sum:", sum, "Difference:", diff)
	fmt.Println(definition(10, 5))
	fmt.Println("--------------------------------------------------")
}
