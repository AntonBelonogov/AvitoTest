package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"AvitoTest/internal/controller"
	"AvitoTest/internal/middleware"
	"AvitoTest/internal/model/entity"
)

func main() {
	router := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(&entity.UserProduct{}, &entity.User{}, &entity.Product{}, &entity.History{}); err != nil {
		log.Fatal(err)
	}

	products := []entity.Product{
		{ID: 1, Name: "t-shirt", Price: 80},
		{ID: 2, Name: "cup", Price: 20},
		{ID: 3, Name: "book", Price: 50},
		{ID: 4, Name: "pen", Price: 10},
		{ID: 5, Name: "powerbank", Price: 200},
		{ID: 6, Name: "hoody", Price: 300},
		{ID: 7, Name: "umbrella", Price: 200},
		{ID: 8, Name: "socks", Price: 10},
		{ID: 9, Name: "wallet", Price: 50},
		{ID: 10, Name: "pink-hoody", Price: 500},
	}

	if dbErr := db.Create(&products).Error; dbErr != nil {
		log.Println("products already exists")
	}

	apiUnprotected := router.Group("/api")
	controller.InitUser(apiUnprotected, db)

	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	controller.InitCoin(api, db)
	controller.InitInfo(api, db)

	if err = router.Run(":" + os.Getenv("SERVER_PORT")); err != nil {
		log.Fatal(err)
	}
}
