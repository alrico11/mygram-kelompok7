package config

import (
	"fmt"
	"os"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	// read db
	dbUsername := os.Getenv("MYSQLUSER")
	dbPassword := os.Getenv("MYSQLPASSWORD")
	dbHost := os.Getenv("MYSQLHOST")
	dbPort := os.Getenv("MYSQLPORT")
	dbName := os.Getenv("MYSQLDATABASE")

	// read db
	dsnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)


	
	db, err := gorm.Open(mysql.Open(dsnString), &gorm.Config{})
	if err != nil {
		fmt.Println(dsnString)
		panic(err.Error())
	}

	return db
}
Footer
© 2022 GitHub, Inc.
