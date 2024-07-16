package main

import (
	"fmt"
	"log"

	"github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/middleware"
	database "github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	err := database.GetDatabaseSingelton().SetDatabase()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("/note/:uuid", database.GetNoteByUuid)
	router.GET("/note", middleware.RequireAuth, database.GetUserNotes)
	router.POST("/note/create", middleware.RequireAuth, database.CreateNote)
	router.POST("/signup", database.SignUp)
	router.POST("/login", database.Login)
	router.GET("/validate", middleware.RequireAuth, database.Validate)
	err = router.Run(":8000")
	if err != nil {
		fmt.Printf("an error occured")
	}
	fmt.Printf("server running on port 8000")
}
