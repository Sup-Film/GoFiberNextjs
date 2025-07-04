package main

import (
	"fmt"
)

// กำหนด Struch Person
type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Go Learning 101 - Chapter 11: Structs")
	fmt.Println("=====================================")

	// สร้าง Instace ของ Struct Person
	person1 := Person{
		Name: "Supakorn",
		Age:  25,
	}

	fmt.Println("Person 1 Name: ", person1.Name)
	fmt.Println("Person 1 Age: ", person1.Age)

	fmt.Println("=====================================")
	fmt.Println("การใช้ Pointer กับ Struct")
	pointerStruch := &Person{
		Name: "John",
		Age:  25,
	}

	fmt.Println("Pointer Person Name: ", pointerStruch.Name)
	fmt.Println("Pointer Person Age: ", pointerStruch.Age)
	editPerson(pointerStruch, "Jane", 30)
	fmt.Println("หลังจากแก้ไขชื่อโดยใช้ Pointer: ", pointerStruch.Name)
	fmt.Println("หลังจากแก้ไขอายุโดยใช้ Pointer: ", pointerStruch.Age)

	fmt.Println("=====================================")
	// การใช้ Method กับ Struct
	fmt.Println("การใช้ Method กับ Struct")
	fmt.Println(person1.SayHello())

	// เรียก Method HaveBirthday
	person1.HaveBirthday()
	fmt.Println("อายุใหม่หลังจากวันเกิด: ", person1.Age)
}

func editPerson(p *Person, name string, age int) {
	p.Name = name
	p.Age = age
}

// สร้าง Function ชื่อ SayHello สำหรับ Struct Person
// * เหมือนกับการสร้าง Class Person ที่มี Method SayHello แล้วมี Constructor เป็น Name และ Age
// * ถ้าเป็น JS จะเป็นแบบนี้
// class Person {
//     constructor(name, age) {
//         this.name = name;
//         this.age = age;
//     }

//	    sayHello() {
//	        return `Hello: ${this.name}, Age: ${this.age}`;
//	    }
//	}
func (p Person) SayHello() string {
	return "Hello: " + p.Name + ", Age: " + fmt.Sprint(p.Age)
}

// Pointer Method
func (p *Person) HaveBirthday() {
	p.Age++
}
