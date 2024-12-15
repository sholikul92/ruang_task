package utils

import (
	"a21hc3NpZ25tZW50/model"
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadConfigDB() (*model.DBConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	dbconfig := model.DBConfig{
		HOST:     os.Getenv("DB_HOST"),
		PORT:     port,
		USER:     os.Getenv("DB_USER"),
		PASSWORD: os.Getenv("DB_PASSWORD"),
		DBNAME:   os.Getenv("DB_NAME"),
	}

	return &dbconfig, nil
}

func LoadToken() (string, error) {
	token := os.Getenv("HUGGINGFACE_TOKEN")
	if token == "" {
		return "", errors.New("HUGGINGFACE_TOKEN is not set in the .env file")
	}

	return token, nil
}
