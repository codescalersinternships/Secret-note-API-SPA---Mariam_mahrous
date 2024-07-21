package main

import (
	"fmt"
	"os"

	app "github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/pkg"
	database "github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/database"

	"github.com/joho/godotenv"
)

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
	a, err := app.NewApp(dbConfig , os.Getenv("TOKEN_KEY") )
	if err != nil {
		fmt.Println("error in config")
	}
	err = a.Run(port , os.Getenv("FRONTEND_URL"))
	if err != nil {
		fmt.Println("error running server")
	}
}
