package main

import (
	"fmt"
)

// การประกาศตัวแปรแบบ Global
var xml string = "outside main" // ประกาศตัวแปร xml เป็น Global Variable

func main() {
	fmt.Println("Variable Types in Go")

	var x int      // ประกาศตัวแปร x เป็นประเภท int
	var y int = 10 // ประกาศตัวแปร y เป็นประเภท int และกำหนดค่าเริ่มต้นเป็น 10
	var z = 20     // ประกาศตัวแปร z โดยไม่ระบุชนิดข้อมูล แต่ Go จะอนุมานชนิดข้อมูลจากค่าเริ่มต้นที่กำหนดให้

	// การประกาศตัวแปรแบบหลายตัว
	var a, b, c int = 1, 2, 3 // ประกาศตัวแปร a, b, c เป็นประเภท int และกำหนดค่าเริ่มต้น
	var (
		a1 int = 5  // ประกาศตัวแปร a1 เป็น int และกำหนดค่าเริ่มต้นเป็น 5
		b1 int = 10 // ประกาศตัวแปร b1 เป็น int และกำหนดค่าเริ่มต้นเป็น 10
	)

	// การประกาศตัวแปรแบบสั้น
	d := 4         // ประกาศตัวแปร d โดยใช้การประกาศแบบสั้น (short variable declaration) และกำหนดค่าเริ่มต้นเป็น 4
	e, f := 20, 30 // ประกาศตัวแปร e, f โดยใช้การประกาศแบบสั้นและกำหนดค่าเริ่มต้น

	// การประกาศตัวแปรแบบค่าคงที่ (constant)
	// การประกาศตัวแปรแบบค่าคงที่จะต้องมีค่าคงที่เสมอกำหนดไว้แล้ว ไม่สามารถเปลี่ยนแปลงค่าได้
	const g = 100             // ประกาศตัวแปร g เป็นค่าคงที่ (constant) และกำหนดค่าเริ่มต้นเป็น 100
	const h = "Hello, World!" // ประกาศตัวแปร h เป็นค่าคงที่ (constant) และกำหนดค่าเริ่มต้นเป็น "Hello, World!"
	const pi = 3.14           // ประกาศตัวแปร pi เป็นค่าคงที่ (constant) และกำหนดค่าเริ่มต้นเป็น 3.14

	fmt.Println("Global Variable xml:", xml) // แสดงค่าของตัวแปร Global
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
	fmt.Println("f:", f)
	fmt.Println("a1:", a1)
	fmt.Println("b1:", b1)
}
