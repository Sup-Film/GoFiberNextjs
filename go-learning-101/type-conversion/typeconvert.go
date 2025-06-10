package main

import (
	"fmt"
	"strconv"
)

// แปลงจาก int เป็น float64
func TypeConvert1() {
	var i int = 42
	var f float64 = float64(i)
	fmt.Println("แปลงจาก int เป็น float")
	fmt.Printf("int: %d To, float: %f\n", i, f)
	fmt.Println("--------------------------------------------------")
}

// แปลงจาก float64 เป็น int
func TypeConvert2() {
	var f float64 = 3.14
	var i int = int(f)
	fmt.Println("แปลงจาก float เป็น int")
	fmt.Printf("float: %f To, int: %d\n", f, i)
	fmt.Println("--------------------------------------------------")
}

// แปลงจาก int เป็น string
func TypeConvert3() {
	var i int = 42
	var s string = strconv.Itoa(i) // ใช้ strconv.Itoa เพื่อแปลง int เป็น string

	fmt.Println("แปลงจาก int เป็น string")
	fmt.Printf("int: %d To, string: %s\n", i, s)
	fmt.Println("--------------------------------------------------")
}

// แปลงจาก String เป็น int พร้อมกับการ Handle error
func TypeConvert4() {
	var s string = "42"
	var i int
	var err error            // เป็นการเก็บ error ที่อาจเกิดขึ้น
	i, err = strconv.Atoi(s) // ผลลัพธ์จะเก็บในตัวแปร i และถ้าเกิดข้อผิดพลาดจะเก็บในตัวแปร err
	// Atoi ย่อมาจาก "ASCII to integer" ซึ่งใช้แปลง string เป็น int

	// การ Handle error ในภาษา Go
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	}

	fmt.Println("แปลงจาก string เป็น int")
	fmt.Printf("string: %s To, int: %d\n", s, i)
	fmt.Println("--------------------------------------------------")
}

// แปลงจาก Bool เป็น int (แสดงในรูปแบบ interger เพราะ Go ไม่มีการแปลงตรงๆ จาก bool เป็น int)
func TypeConvert5() {
	var b bool = true
	var i int

	if b {
		i = 1 // ถ้า b เป็น true ให้ i เป็น 1
	} else {
		i = 0 // ถ้า b เป็น false ให้ i เป็น 0
	}
	fmt.Println("แปลงจาก bool เป็น int โดยใช้เงื่อนไข if-else (แสดงในรูปแบบ integer)")
	fmt.Printf("bool: %t To, int: %d\n", b, i)
	fmt.Println("--------------------------------------------------")
}
