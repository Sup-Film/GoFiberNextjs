package models

// กำหนด Struct สำหรับ User
type User struct {
	ID        int    `json:"id"`         // ID ของผู้ใช้
	Username  string `json:"username"`   // ชื่อผู้ใช้
	Password  string `json:"password"`   // รหัสผ่าน
	Email     string `json:"email"`      // อีเมลของผู้ใช้
	Fullname  string `json:"fullname"`   // ชื่อเต็มของผู้ใช้
	Role      string `json:"role"`       // บทบาทของผู้ใช้ (เช่น admin, user)
	CreatedAt string `json:"created_at"` // วันที่สร้างผู้ใช้
	UpdatedAt string `json:"updated_at"` // วันที่ปรับปรุงข้อมูลผู้ใช้
}
