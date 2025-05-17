package database

import (
	"fmt"
	"log"
	"os"

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
	host := getEnv("DB_HOST")
	user := getEnv("DB_USER")
	password := getEnv("DB_PASSWORD")
	dbname := getEnv("DB_NAME")
	port := getEnv("DB_PORT")

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

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Println("Environment variable not found:", key)
	return os.DevNull
}
