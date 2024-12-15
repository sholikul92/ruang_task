package config

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnctPostgreSql(dbConfig *model.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", dbConfig.HOST, dbConfig.USER, dbConfig.PASSWORD, dbConfig.DBNAME, dbConfig.PORT)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	return db, nil
}
