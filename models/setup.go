package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

//connection user
func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(192.168.138.139:3306)/kammi"))
	if err != nil {
		fmt.Println("Gagal koneksi database")
	}

	DB = db
}

//192.168.138.139
