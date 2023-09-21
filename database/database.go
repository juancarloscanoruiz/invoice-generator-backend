package database

import (
	"invoice-app/database/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	dsn := "root:museshowbiz@tcp(127.0.0.1:3306)/invoice_generator?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	db.AutoMigrate(&models.Invoice{}, &models.Client{}, &models.Item{})
}
