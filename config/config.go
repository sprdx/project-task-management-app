package config

import (
	"task-management/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Define GORM Database
var DB *gorm.DB

// Define global error variable for handling error
var err error

//Declare function to connect database
func InitDB() {
	//Set data source that will be used
	connection := "root:qwerty@tcp(127.0.0.1:3306)/todo_list?charset=utf8mb4&parseTime=True&loc=Local"

	//Initialize DB session
	DB, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	//Migrate the database schema
	InitMigration()
}

//Declare function to auto-migrate the schema
func InitMigration() {
	DB.AutoMigrate(&models.User{}, models.Task{}, models.Project{})
}
