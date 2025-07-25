package config

import (
	"log"
	"os"

	"coop_student_backend/internal/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	connectionString := os.Getenv("DATABASE_URL")
	db, error := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if error != nil {
		log.Fatal("Fail to connect to DB : ",error)
	}
	db.AutoMigrate(&domain.UserLogin{})
	db.AutoMigrate(&domain.Payment{})
	db.AutoMigrate(&domain.Employee{})
	db.AutoMigrate(&domain.Product{})
	
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Order{})
	db.AutoMigrate(&domain.OrderDetail{})
	return db
}