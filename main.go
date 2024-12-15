package main

import (
	"log"
	"os"

	db "a21hc3NpZ25tZW50/config"
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/router"
	"a21hc3NpZ25tZW50/utils"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token, err := utils.LoadToken()
	if err != nil {
		log.Fatal(err)
	}

	dbConfig, err := utils.LoadConfigDB()
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.ConnctPostgreSql(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal(err)
	}

	client := resty.New()
	base_url := "https://api-inference.huggingface.co/models/"

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := router.SetupRouter(client, token, base_url, db)

	log.Fatal(r.Run(":" + port))
}
