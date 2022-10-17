package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"cutloss-trading/app/models"
)

var db *gorm.DB
var err error

func InitStore() (*gorm.DB, error) {
	loadEnv()
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	fmt.Println("Password db: ", os.Getenv("POSTGRES_PASSWORD"))
	db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	migrateError := db.AutoMigrate(&models.User{})

	if migrateError != nil {
		log.Fatalln(migrateError)
	}

	return db, nil
}

func loadEnv() {
	env := os.Getenv("ENV")
	if "" == env {
		env = "development"
	}
	envLoad := godotenv.Load(".env." + env + ".local")
	// var envLoad2 error
	// if "test" != env {
	// 	envLoad2 = godotenv.Load(".env.local")
	// }
	// envLoad3 := godotenv.Load(".env." + env)
	envLoad4 := godotenv.Load() // The Original .env

	if envLoad != nil {
		log.Fatalln(envLoad)
	}

	// if envLoad2 != nil {
	// 	log.Fatalln(envLoad2)
	// }

	// if envLoad3 != nil {
	// 	log.Fatalln(envLoad3)
	// }

	if envLoad4 != nil {
		log.Fatalln(envLoad4)
	}
}

func DbManager() *gorm.DB {
	return db
}
