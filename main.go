package main

import (
	"fmt"
	"log"
	"os"
	"statbot/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectPostgres() *gorm.DB {
	// connect to postgres using gorm
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&models.Users{}, &models.Word{})
	return db
}

func main() {
	dB := connectPostgres()

	log.Println(dB.DB())
}
