package main

import "fmt"

func SampleArray() {

	// * ประกาศ Array ว่าง ๆ
	var fruits [3]string

	// กำหนดค่าให้กับ Array
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Cherry"

	// * การเช้าถึงแบบใช้ index
	fmt.Println("การเข้าถึงค่าใน Array โดยใช้ index:")
	fmt.Println("Fruits in the array:", fruits[0], fruits[1], fruits[2])

	// * การวนลูปเพื่อแสดงผลค่าใน Array
	fmt.Println("การเข้าถึงค่าใน Array โดยใช้ for loop:")
	for i, fruit := range fruits {
		fmt.Printf("Fruit at index %d: %s\n", i, fruit)
	}
	fmt.Println("--------------------------------------------------")

	// * นับจำนวนสมาชิกใน Array
	fmt.Println("จำนวนสมาชิกใน Array:", len(fruits))
	fmt.Println("--------------------------------------------------")

}

// * การสร้าง Array และกำหนดสมาชิก
func SampleArray2() {
	// * ประกาศ Array พร้อมกำหนดค่าเริ่มต้น
	var carbrands = [5]string{"Honda", "Toyota", "Honda", "Mazda", "BMW"}
	fmt.Println("สมาชิกใน Array carbrands:", carbrands) // [Honda Toyota Honda Mazda BMW]

	fmt.Println("--------------------------------------------------")

	// * กรณีกำหนดสมาชิก Array ไม่ครบ
	var numbers = [5]int{1, 2, 3}                     // จะมีสมาชิกที่เหลือเป็นค่าเริ่มต้น (0)
	fmt.Println("กำหนดสมาชิก Array ไม่ครบ:", numbers) // [1 2 3 0 0]

	fmt.Println("--------------------------------------------------")

	// * กำหนดตำแหน่งสมาชิกใน Array
	var colors = [10]int{0: 1, 4: 2, 6: 3}                    // กำหนดสมาชิกที่ตำแหน่ง 0, 4, และ 6
	fmt.Println("กำหนดสมาชิก Array ที่ตำแหน่งเฉพาะ:", colors) // [1 0 0 0 2 0 3 0 0 0]

	fmt.Println("--------------------------------------------------")

	// * Array แบบไม่จำกัดขนาดสมาชิก
	// * แต่ไม่สามารถเพิ่มสมาชิกได้ภายหลัง
	var array1 = [...]int{1, 2, 3, 4, 5}                // ใช้ ... เพื่อให้ Go คำนวณขนาดของ Array อัตโนมัติ
	fmt.Println("Array แบบไม่จำกัดขนาดสมาชิก:", array1) // [1 2 3 4 5]

}

func ArrayMultiDimensional() {
	// * ประกาศ Array แบบหลายมืติ พร้อมกำหนดค่าเริ่มต้น
	var multiDimensional = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	for i, row := range multiDimensional {
		for j, valie := range row {
			fmt.Printf("Value at [%d][%d]: %d\n", i, j, valie)
		}
	}
	fmt.Println("--------------------------------------------------")

	var matrix = [2][2]int{{1, 2}, {3, 4}}

	for i, row := range matrix {
		fmt.Printf("Row %d:\n", i)
		for j, value := range row {
			fmt.Printf("Matrix[%d][%d]: %d\n", i, j, value)
		}
	}
}
