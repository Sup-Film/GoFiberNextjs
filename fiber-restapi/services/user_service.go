package services

import (
	"fiber-restapi/models" // อ้างอิงไปยังโมเดล
	"fmt"
	_ "fmt"
	"time"
	_ "time"
)

// UserService struct สำหรับจัดการผู้ใช้
type UserService struct {
	users  []models.User // จำลองฐานข้อมูลผู้ใช้ โดยใช้ slice ของ User | เป็นตัวแทนของฐ้านข้อมูล ซึ่งเราจำลองเอาไว้ในหน่วยความจำ
	nextID int           // ตัวนับ ID ถัดไปสำหรับผู้ใช้ใหม่ | ใช้สำหรับการจำลองฐานข้อมูล
}

// สร้างฟังก์ชั่นสำหรับ NewUserService สร้างผู้ใช้ใหม่
func NewUserService() *UserService {
	return &UserService{
		users:  []models.User{}, // เริ่มต้นด้วย slice ว่าง
		nextID: 1,               // เริ่มต้น ID ถัดไปที่ 1
	}
}

// สร้างฟังก์ชั่นสำหรับ RegisterUser ลงทะเบียนผู้ใช้ใหม่
// (s *UserService) Method นี้เป็นของ UserService struct
// (RegisterUser(user models.User)) รับข้อมูล User ที่จะลงทะเบียน
// (*models.User, error) คืนค่าผู้ใช้ที่ลงทะเบียนใหม่และข้อผิดพลาด (ถ้ามี)
func (s *UserService) RegisterUser(user models.User) (*models.User, error) {
	// ตรวจสอบว่ามีผู้ใช้อยู่แล้วหรือไม่
	for _, u := range s.users {
		if u.Username == user.Username {
			// (*models.User) nil = ไม่มี User ที่สร้างสำเร็จ (เพราะเจอ duplicate)
			// (error) fmt.Errorf(...) = error message อธิบายว่าทำไมไม่สำเร็จ
			return nil, fmt.Errorf("username %s already exists", user.Username)
		}
		if u.Email == user.Email {
			return nil, fmt.Errorf("email %s already exists", user.Email)
		}
	}

	// กำหนด ID และ TimeStamp ให้กับผู้ใช้ใหม่
	user.ID = s.nextID
	s.nextID++                             // เพิ่ม ID ถัดไป
	now := time.Now().Format(time.RFC3339) // ใช้รูปแบบ ISO 8601 สำหรับวันที่และเวลา
	user.CreatedAt = now
	user.UpdatedAt = now

	// กำหนดค่าเริ่มต้นให้กับ Role
	if user.Role == "" {
		user.Role = "user" // กำหนด Role เป็น "user" ถ้าไม่ได้ระบุ
	}

	// เพิ่มผู้ใช้ใหม่เข้าไปใน slice
	s.users = append(s.users, user)
	return &user, nil // คืนค่าผู้ใช้ที่ลงทะเบียนใหม่และไม่มีข้อผิดพลาด
}

// สร้างฟังก์ชั่นสำหรับ Login เข้าสู่ระบบผู้ใช้
func (s *UserService) Login(username, password string) (*models.User, string, error) {
	for _, u := range s.users {
		if u.Username == username && u.Password == password {
			// สร้าง Token แบบง่าย (การใช้งานจริงต้องใช้ JWT)
			token := fmt.Sprintf("token_%s_%d", u.Username, time.Now().Unix())
			return &u, token, nil // คืนค่าผู้ใช้และ Token
		}
	}

	return nil, "", fmt.Errorf("invalid username or password")
}
