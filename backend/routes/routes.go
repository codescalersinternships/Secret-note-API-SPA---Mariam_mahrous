package routes

import (
	"github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/middleware"
	database "github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/pkg"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/note/:uuid", database.GetNoteByUuid)
	router.GET("/note", middleware.RequireAuth, database.GetUserNotes)
	router.POST("/note/create", middleware.RequireAuth, database.CreateNote)
	router.POST("/signup", database.SignUp)
	router.POST("/login", database.Login)
	router.GET("/validate", middleware.RequireAuth, database.Validate)
}