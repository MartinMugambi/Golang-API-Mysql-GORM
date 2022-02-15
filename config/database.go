package config

import (
	"backend/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var Database *gorm.DB

var err error

const dns = "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

func InitializeDatabase() {
	Database, err = gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Could not connect to Database")
	}

	databaseUser := os.Getenv("DATABASE_USER")

	fmt.Println(databaseUser)

	Database.AutoMigrate(&models.User{})

}
