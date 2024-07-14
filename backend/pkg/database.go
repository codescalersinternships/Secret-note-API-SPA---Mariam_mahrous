package db

import (
	"github.com/joho/godotenv"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
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

func (s * DatabaseSingleton) SetDatabase() {
	err := godotenv.Load(".env")
	if err != nil {
		 panic(err.Error()) 
	}

	db := os.Getenv("DATABASE")
	if db =="POSTGRESQL"{
		dsn := "host=localhost user=postgres password=123456 dbname=postgres port=5432"
		s.database, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		s.database, _= gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	}
}