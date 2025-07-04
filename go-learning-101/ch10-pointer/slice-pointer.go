package main

import (
	"fmt"
)

func PointerSlice() {
	// Create a slice of integers
	slice := []int{1, 2, 3, 4, 5}

	// * สร้าง Pointer ชี้ไปที่ slice
	pointerSlice := &slice

	fmt.Println("Slice:", slice)
	fmt.Println("Pointer to Slice:", pointerSlice)
	fmt.Println("ค่าในของ Value ที่ Pointer ชี้ไป:", *pointerSlice)

	// * เปลี่ยนค่าใน slice ผ่าน pointer
	(*pointerSlice)[0] = 10

	fmt.Println("Slice หลังจากเปลี่ยนค่าผ่าน Pointer:", slice)

	// * วนลูปผ่าน pointer เพื่อแสดงค่าทั้งหมดใน slice
	for i, v := range *pointerSlice {
		fmt.Printf("Index %d: Value %d\n", i, v)
	}
}
