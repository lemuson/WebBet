package backend

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	string_connection := "root:0000@tcp(127.0.0.1:3306)/sportsbetting?charset=utf8mb4"

	Test_DataBase, err := gorm.Open(mysql.Open(string_connection), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	return Test_DataBase
}
