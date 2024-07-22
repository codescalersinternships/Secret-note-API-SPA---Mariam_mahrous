package main

import (
	"fmt"
	"os"

	database "github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/database"
	_ "github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/docs"
	app "github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/pkg"
	"github.com/joho/godotenv"
)

// @title Secret-Note API/SPA
// @version 1.0
// @description A secret note api in go using Gin framework
// @host localhost:8080
func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading .env")
	}
	port := os.Getenv("PORT")
	dbConfig := database.ConfigDB{
		DatabaseType: os.Getenv("DATABASE"),
		Host:         "localhost",
		User:         os.Getenv("USER"),
		Password:     os.Getenv("PASSWORD"),
		DBName:       os.Getenv("DBNAME"),
		Port:         os.Getenv("POSTGRESQL_PORT"),
	}
	a, err := app.NewApp(dbConfig, os.Getenv("TOKEN_KEY"))
	if err != nil {
		fmt.Println("error in config")
	}
	err = a.Run(port, os.Getenv("FRONTEND_URL"))
	if err != nil {
		fmt.Println("error running server")
	}
}
