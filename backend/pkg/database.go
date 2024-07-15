package db

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseSingleton struct {
	database *gorm.DB
}

var database *DatabaseSingleton

func GetDatabaseSingelton() *DatabaseSingleton {
	if database == nil {
		database = &DatabaseSingleton{database: nil}
	}
	return database
}

func (s *DatabaseSingleton) GetDatabase() *gorm.DB {
	return s.database
}

func (s *DatabaseSingleton) SetDatabase() error {
	err := godotenv.Load(".env")
	if err != nil {
		return errors.New("error loading .env")
	}

	db := os.Getenv("DATABASE")
	var dbErr error
	if db == "POSTGRESQL" {
		dsn := "host=localhost user=postgres password=123456 dbname=postgres port=5432"
		s.database, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if dbErr != nil {
			return fmt.Errorf("failed to connect to PostgreSQL database: %v", dbErr.Error())
		}
	} else {
		s.database, dbErr = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if dbErr != nil {
			return fmt.Errorf("failed to connect to Sqlite database: %v", dbErr.Error())
		}
	}
	dbErr = database.database.AutoMigrate(&Note{}, &User{})
	if dbErr != nil {
		return fmt.Errorf("failed to migrate to database: %v", dbErr.Error())
	}
	return nil
}
