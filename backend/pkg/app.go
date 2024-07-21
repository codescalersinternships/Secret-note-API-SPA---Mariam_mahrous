package app

import (
	"fmt"
	"net/http"

	model "github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/models"
	middleware "github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/middleware"
	database "github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type App struct {
	R  *gin.Engine
	DB database.DB
	MW middleware.MW 
}

func NewApp(db database.ConfigDB , tokenKey string) (*App, error) {
	database, err := database.NewDB(db)
	mw := middleware.MW{
		TokenKey: tokenKey,
		DB:       *database,
	}
	if err != nil {
		return &App{R: gin.Default(), DB: *database , MW: mw}, err
	}
	return &App{R: gin.Default(), DB: *database , MW: mw}, nil
}

func (a *App) RegisterHandlers() {
	a.R.GET("/note/:uuid", a.GetNoteByUuid)
	a.R.GET("/note", a.MW.RequireAuth, a.GetUserNotes)
	a.R.POST("/note/create", a.MW.RequireAuth, a.CreateNote)
	a.R.POST("/signup", a.SignUp)
	a.R.POST("/login", a.Login)
}

func (a *App) Run(port string , frontendUrl string) error {

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{frontendUrl}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true

	a.R.Use(cors.New(config))

	a.RegisterHandlers()
	err := a.R.Run(":" + port)
	if err != nil {
		return fmt.Errorf("failed to run server on port: %s", port)
	}
	return nil
}

func (a *App) GetNoteByUuid(c *gin.Context) {
	database := a.DB
	uuid, _ := uuid.Parse(c.Param("uuid"))
	note, statusCode, err := database.GetNoteByUuid(uuid)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(statusCode, note)
}

func (a *App) GetUserNotes(c *gin.Context) {
	id, ok := c.Get("id")
	fmt.Print("notes 1")
	if !ok {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		fmt.Print("notes 2")
		return
	}
	notes, statusCode, err := a.DB.GetAllNotes(id.(uint))
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
		fmt.Print("notes 3")
		return
	}
	c.IndentedJSON(statusCode, notes)
}

func (a *App) CreateNote(c *gin.Context) {
	var newNote model.Note
	if err := c.BindJSON(&newNote); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Check that all fields have been sent"})
		return
	}
	id, ok := c.Get("id")
	if !ok {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}
	note, statusCode, err := a.DB.CreateNote(id.(uint), newNote)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(statusCode, note)
}

func (a *App) SignUp(c *gin.Context) {
	var newUser model.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	hashedpassword , statusCode ,err := a.MW.HashPassword(newUser.Password)
	if err !=nil{
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	newUser.Password = hashedpassword

	tokenString, statusCode, err := a.MW.GenerateToken(newUser)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	newUser.Token = tokenString
	
	statusCode, err = a.DB.SignUp(newUser)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
 		return
	}
	c.IndentedJSON(statusCode, newUser)
}


func (a *App) Login(c *gin.Context) {
	var newUser model.User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	registeredUser, found := a.DB.GetUser(newUser.Email)
	if !found {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Invalid email or password"})
		return
	}

	statusCode , err:= a.MW.CompareHashAndPassword(registeredUser.Password, newUser.Password )
	if err!= nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
	}
	tokenString, statusCode, err := a.MW.GenerateToken(registeredUser)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
	}
	registeredUser.Token = tokenString
	//save user to db
	c.IndentedJSON(http.StatusOK, registeredUser)
}

