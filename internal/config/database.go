package config

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func ConnectDb() *gorm.DB{


	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("file .env tidak di temukan")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := "127.0.0.1"
	dbPort := "5433"

dsn := fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", 
    dbHost, dbUser, dbPass, dbName, dbPort,
)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("gagal terhubung ke database: %s\n", err)

	}
	fmt.Println("berhasil terhubung ke database menggunakan gorm")
	DB := db
	return DB
}