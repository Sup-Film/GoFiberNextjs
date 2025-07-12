package db

import (
	"fmt"
	"log"

	"github.com/Sup-Film/fiber-nextjs-ecommerce/config"
	"github.com/Sup-Film/fiber-nextjs-ecommerce/internal/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect เชื่่อมต่อกับฐานข้อมูล PostgreSQL
func Connect(config *config.Config) *gorm.DB {
	// สร้าง Connection String
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Bangkok", config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort, config.DBSSL)

	// เชื่อมต่อกับฐานข้อมูล PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Migrate ตารางต่างๆ ในฐานข้อมูล
	migrateDatabase(db)
	log.Println("Connected to PostgreSQL database")

	return db
}

// Migrate ทำการ Migrate ตารางต่างๆ ในฐานข้อมูล
func migrateDatabase(db *gorm.DB) {
	// Auto Migrate สร้างตารางในฐานข้อมูลตามโครงสร้างของโมเดล
	err := db.AutoMigrate(
		&domain.Role{},
		&domain.Permission{},
		&domain.User{},
		&domain.Category{},
		&domain.Product{},
		&domain.ProductImage{},
		&domain.Cart{},
		&domain.CartItem{},
		&domain.Order{},
		&domain.OrderItem{},
		&domain.Transaction{},
	)

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database migration completed successfully")
}
