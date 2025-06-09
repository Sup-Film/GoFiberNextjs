package greeting

import "fmt"

func SayMeeting() {
	fmt.Println("Hello, From Meeting method of Greeting Package")
}

// เมื่อประกาศฟังก์ชันใน package เป็นตัวพิมพ์เล็ก จะทำให้ฟังก์ชันนั้นเป็น private
// และไม่สามารถเข้าถึงได้จาก package อื่น
// หากต้องการให้ฟังก์ชันนี้สามารถเข้าถึงได้จาก package อื่น ต้องเปลี่ยนชื่อฟังก์ชันให้เป็นตัวพิมพ์ใหญ่
// เช่น เปลี่ยนจาก sayhi เป็น SayHi
func sayhi() {
	fmt.Println("Hello, From private method of Greeting Package")
}
