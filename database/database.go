package database

import (
	"fmt"
	"log"

	"hello-go/utils"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using default values")
	}
	host := utils.GetEnv("DB_HOST", "localhost")
	user := utils.GetEnv("DB_USER", "postgres")
	password := utils.GetEnv("DB_PASSWORD", "password")
	dbname := utils.GetEnv("DB_NAME", "mydb")
	port := utils.GetEnv("DB_PORT", "5432")

	// Connection URL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, dbname, port)

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Connected to database")
	DB = db
}
