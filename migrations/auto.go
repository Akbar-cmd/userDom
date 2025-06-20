package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"

	"userDomain/internal/user"
)

func main() {

	log.Println("Запускаем миграцию...")
	err := godotenv.Load("config/local.yaml")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("db_dsn")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&user.User{})
	log.Println("Миграция завершена...")
}
