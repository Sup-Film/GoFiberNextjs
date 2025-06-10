package main

import (
	"fmt"
)

func Forloop101() {
	for i := 0; i < 3; i++ {
		fmt.Print(i, ", ")
	}
}

func Forwhile() {
	sum := 10

	for sum <= 20 {
		sum += 1
		fmt.Print(sum, ", ")
	}
}

func ForDoWhile() {
	sum := 1

	for {
		if sum >= 5 {
			break // ออกจากลูปเมื่อ sum >= 5
		}

		fmt.Print(sum, ", ")
		sum += 1
	}
}
