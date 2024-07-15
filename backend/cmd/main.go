package main

import (
	"fmt"
	"log"

	database "github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	err := database.GetDatabaseSingelton().SetDatabase()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("./note/:uuid", database.GetNoteByUuid)
	router.POST("./note/create", database.CreateNote)
	err = router.Run(":8000")
	if err != nil {
		fmt.Printf("an error occured")
	}
	fmt.Printf("server running on port 8000")
}
