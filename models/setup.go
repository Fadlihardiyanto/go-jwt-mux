package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_jwt_mux"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connect database: ", err)
	}

	db.AutoMigrate(&User{})
	DB = db

}
