package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

func ConectDB() *gorm.DB {
	var db *gorm.DB

	getEnvVar := godotenv.Load()
	if getEnvVar != nil {
		log.Printf("do not .env variables: %v", getEnvVar)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
					   os.Getenv("DB_USER"),
				       os.Getenv("DB_PASSWORD"),
					   os.Getenv("DB_HOST"),
					   os.Getenv("DB_PORT"),
					   os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error to conect database %v", err)
	}
	log.Println("database conected")
	return db
}
