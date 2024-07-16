package main

import (
	"fmt"
	"os"

	app "github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/pkg"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading .env")
	}
	port := os.Getenv("PORT")
	dbConfig := app.ConfigDB{
		DatabaseType: os.Getenv("DATABASE"),
		Host:         "localhost",
		User:         os.Getenv("USER"),
		Password:     os.Getenv("PASSWORD"),
		DBName:       os.Getenv("DBNAME"),
		Port:         os.Getenv("POSTGRESQL_PORT"),
	}
	a, err := app.NewApp(dbConfig)
	if err != nil {
		fmt.Println("error in config")
	}
	err = a.Run(port)
	if err != nil {
		fmt.Println("error running server")
	}
}
