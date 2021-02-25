package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	viperUser := os.Getenv("PGNAME")
	viperPassword := os.Getenv("PGPASSWORD")
	viperDb := os.Getenv("PGDATABASE")
	viperHost := os.Getenv("PGHOST")
	viperPort := os.Getenv("PGPORT")

	prosgretConname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", viperHost, viperPort, viperUser, viperDb, viperPassword)
	fmt.Println("conname is\t\t", prosgretConname)
	db, err := gorm.Open(postgres.Open(prosgretConname), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db

}
