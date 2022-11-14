package config

import (
	"fmt"
	"log"
	"os"
	"project2/model/entity"
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
		if err != nil {
		log.Fatal("DB Konek Eror")
	}
	fmt.Println("DB Berhasil Konek")
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Comment{})
	db.AutoMigrate(&entity.Photo{})
	db.AutoMigrate(&entity.SocialMedia{})

	return db
}
