package config

import (
	"fmt"
	"os"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	// read db
	dbUsername := os.Getenv("PGUSER")
	dbPassword := os.Getenv("PGPASSWORD")
	dbHost := os.Getenv("PGHOST")
	dbPort := os.Getenv("PGPORT")
	dbName := os.Getenv("PGDATABASE")

	// read db
	dsnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)


	
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(dsnString)
		panic(err.Error())
	}

	return db
}
