package main

import (
	"fmt"
)

func main() {
	fmt.Println("Operators in Go")

	var x int

	// การบวก
	x = 10 + 5 // x จะมีค่าเป็น 15
	fmt.Println("10 + 5 =", x)

	// การลบ
	x = 10 - 5 // x จะมีค่าเป็น 5
	fmt.Println("10 - 5 =", x)

	// การคูณ
	x = 10 * 5 // x จะมีค่าเป็น 50
	fmt.Println("10 * 5 =", x)

	// การหาร
	x = 10 / 5 // x จะมีค่าเป็น 2
	fmt.Println("10 / 5 =", x)

	// การหารเศษ
	x = 10 % 3 // x จะมีค่าเป็น 1 (เศษจากการหาร 10 ด้วย 3)
	fmt.Println("10 % 3 =", x)

	// การเพิ่มค่า
	x = 10
	x++ // x จะมีค่าเป็น 11
	fmt.Println("10++ =", x)

	// การลดค่า
	x-- // x จะมีค่าเป็น 10
	fmt.Println("11-- =", x)

	// การเปรียบเทียบ
	a := 10
	b := 20
	fmt.Println("10 == 20:", a == b) // false
	fmt.Println("10 != 20:", a != b) // true
	fmt.Println("10 < 20:", a < b)   // true
	fmt.Println("10 <= 20:", a <= b) // true
	fmt.Println("10 > 20:", a > b)   // false
	fmt.Println("10 >= 20:", a >= b) // false

	// การดำเนินการทางตรรกะ
	fmt.Println("true && false:", true && false) // false
	fmt.Println("true || false:", true || false) // true
	fmt.Println("!true:", !true)                 // false

	// การดำเนินการทางบิต
	x = 5                         // 0101
	y := 3                        // 0011
	fmt.Println("5 & 3 =", x&y)   // AND: 0101 & 0011 = 0001 (1)
	fmt.Println("5 | 3 =", x|y)   // OR: 0101 | 0011 = 0111 (7)
	fmt.Println("5 ^ 3 =", x^y)   // XOR: 0101 ^ 0011 = 0110 (6)
	fmt.Println("5 << 1 =", x<<1) // Shift Left: 0101 << 1 = 1010 (10)
	fmt.Println("5 >> 1 =", x>>1) // Shift Right: 0101 >> 1 = 0010 (2)
	fmt.Println("5 &^ 3 =", x&^y) // AND NOT: 0101 &^ 0011 = 0100 (4)
	fmt.Println("5 << 2 =", x<<2) // Shift Left: 0101 << 2 = 10100 (20)
	fmt.Println("5 >> 2 =", x>>2) // Shift Right: 0101 >> 2 = 0001 (1)
	fmt.Println("5 &^ 3 =", x&^y) // AND NOT: 0101 &^ 0011 = 0100 (4)
}
