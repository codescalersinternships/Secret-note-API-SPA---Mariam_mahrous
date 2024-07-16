package main

import (
	"fmt"
	"os"

	"github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/routes"

	database "github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/pkg"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)




func main() {

	err := godotenv.Load(".env")
	if err != nil {
		//ADD SLOG HENA
		fmt.Println("error loading .env")
	}
	port := os.Getenv("PORT")
	//databaseType := os.Getenv("DATABASE")


	err = database.GetDatabaseSingelton().SetDatabase()
	if err != nil {
		fmt.Println("error accessing database")
	}
	

	router := gin.Default()
	routes.Routes(router)

	err = router.Run(":" + port)
	if err != nil {
		fmt.Println("an error occured connecting to the database")
	}
	fmt.Printf("server running on port: %v" , port)
}
