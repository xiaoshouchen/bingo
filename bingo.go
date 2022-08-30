package main

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"os"
)

var Db *gorm.DB

func main() {
	dsn := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dsn = fmt.Sprintf(dsn, user, password, host, port)
	_, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	} else {

	}

}
