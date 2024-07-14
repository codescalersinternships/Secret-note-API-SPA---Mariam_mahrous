package main

import (
	note "github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/backend/pkg"
	"github.com/joho/godotenv"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"fmt"
	"github.com/google/uuid"
)



func main() {
	var database *gorm.DB
	err := godotenv.Load(".env")
	if err != nil {
		 panic(err.Error()) 
	}

	db := os.Getenv("DATABASE")
	if db =="POSTGRESQL"{
		dsn := "host=localhost user=postgres password=123456 dbname=postgres port=5432"
		database, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		database, _= gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	}

	database.AutoMigrate(&note.Note{})
	database.AutoMigrate(&note.Note{})
	
	newnote := note.CreateNote("mimo", "content",  "2024-07-17" , 10,database)
	fmt.Println(newnote)
	uuid , _ := uuid.Parse("40c0d3ad-0d16-4973-95bb-a1624b725bc6")
	targetNote := note.GetNoteByUuid(uuid,database)
	fmt.Println(targetNote)
}