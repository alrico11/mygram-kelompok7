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
	 psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    dbHost, dbPort, dbUsername, dbPassword, dbName)


	
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(dsnString)
		panic(err.Error())
	}

	return db
}
