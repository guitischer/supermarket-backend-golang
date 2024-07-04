package database

import (
	"fmt"
	"os"

	models "supermarket-api/src/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	err := godotenv.Load("./../.env")
	if err != nil {
		panic("Falha ao carregar arquivo .env")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s"+
		" sslmode=disable TimeZone=America/Sao_Paulo",
		os.Getenv("POSTGRESQL_HOST"), os.Getenv("POSTGRESQL_USER"),
		os.Getenv("POSTGRESQL_PASSWORD"), os.Getenv("POSTGRESQL_DATABASE"),
		os.Getenv("POSTGRESQL_PORT"),
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Falha ao conectar com o banco de dados")
	}

	err = database.AutoMigrate(&models.ProductType{}, &models.Product{})
	if err != nil {
		return
	}

	DB = database
}
