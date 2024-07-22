package app

import (
	"fmt"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	database "github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/database"
	middleware "github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/middleware"
	model "github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type App struct {
	R  *gin.Engine
	DB database.DB
	MW middleware.MW
}

func NewApp(db database.ConfigDB, tokenKey string) (*App, error) {
	database, err := database.NewDB(db)
	mw := middleware.MW{
		TokenKey: tokenKey,
		DB:       *database,
	}
	if err != nil {
		return &App{R: gin.Default(), DB: *database, MW: mw}, err
	}
	return &App{R: gin.Default(), DB: *database, MW: mw}, nil
}

func (a *App) RegisterHandlers() {
	a.R.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	a.R.GET("/note/:uuid", a.GetNoteByUuid)
	a.R.GET("/note", a.MW.RequireAuth, a.GetUserNotes)
	a.R.POST("/note/create", a.MW.RequireAuth, a.CreateNote)
	a.R.POST("/signup", a.SignUp)
	a.R.POST("/login", a.Login)
}

func (a *App) Run(port string, frontendUrl string) error {

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

// @Summary Get note by UUID
// @Description Retrieve a note by its UUID.
// @Tags notes
// @Accept json
// @Produce json
// @Param uuid path string true "Note UUID"
// @Success 200 {object} model.Note "Successful operation"
// @Failure 400 {object} model.ErrorResponse "Invalid UUID format"
// @Failure 404 {object} model.ErrorResponse "Note not found"
// @Router /note/{uuid} [get]
func (a *App) GetNoteByUuid(c *gin.Context) {
	database := a.DB
	uuid, _ := uuid.Parse(c.Param("uuid"))
	note, statusCode, err := database.GetNoteByUuid(uuid)
	if err != nil {
		c.IndentedJSON(statusCode, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.IndentedJSON(statusCode, note)
}

// @Summary Get all notes for the authenticated user
// @Description Retrieve all notes for the authenticated user.
// @Tags notes
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {array} model.Note "Successful operation"
// @Failure 401 {object} model.ErrorResponse "Unauthorized access"
// @Router /note [get]
func (a *App) GetUserNotes(c *gin.Context) {
	id, ok := c.Get("id")
	if !ok {
		c.IndentedJSON(http.StatusUnauthorized, model.ErrorResponse{Error: "Unauthorized access"})
		return
	}
	notes, statusCode, err := a.DB.GetAllNotes(id.(uint))
	if err != nil {
		c.IndentedJSON(statusCode, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.IndentedJSON(statusCode, notes)
}

// @Summary Create a new note
// @Description Create a new note for the authenticated user.
// @Tags notes
// @Accept json
// @Produce json
// @Param note body model.CreateNoteInput true "Note data"
// @Param Authorization header string true "Bearer token"
// @Success 201 {object} model.Note "Note created successfully"
// @Failure 400 {object} model.ErrorResponse "Invalid input"
// @Failure 401 {object} model.ErrorResponse "Unauthorized access"
// @Router /note/create [post]
func (a *App) CreateNote(c *gin.Context) {
	var newNote model.Note
	if err := c.BindJSON(&newNote); err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "Check that all fields have been sent"})
		return
	}
	id, ok := c.Get("id")
	if !ok {
		c.IndentedJSON(http.StatusUnauthorized, model.ErrorResponse{Error: "Unauthorized access"})
		return
	}
	note, statusCode, err := a.DB.CreateNote(id.(uint), newNote)
	if err != nil {
		c.IndentedJSON(statusCode, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, note)
}

// @Summary Sign up a new user
// @Description Sign Up --- Create a new user account.
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.CreateUserInput true "User data"
// @Success 201 {object} model.User "User created successfully"
// @Failure 400 {object} model.ErrorResponse "Invalid input"
// @Failure 409 {object} model.ErrorResponse "Can't create user"
// @Router /signup [post]
func (a *App) SignUp(c *gin.Context) {
	var newUser model.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "Bad request"})
		return
	}

	hashedpassword, statusCode, err := a.MW.HashPassword(newUser.Password)
	if err != nil {
		c.IndentedJSON(statusCode, model.ErrorResponse{Error: err.Error()})
		return
	}
	newUser.Password = hashedpassword

	tokenString, statusCode, err := a.MW.GenerateToken(newUser)
	if err != nil {
		c.IndentedJSON(statusCode, model.ErrorResponse{Error: err.Error()})
		return
	}
	newUser.Token = tokenString

	statusCode, err = a.DB.SignUp(newUser)
	if err != nil {
		c.IndentedJSON(statusCode, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.IndentedJSON(statusCode, newUser)
}

// @Summary Log in a user
// @Description Login --Authenticate a user and return a token.
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.CreateUserInput true "User credentials"
// @Success 200 {object} model.User "Successful login"
// @Failure 400 {object} model.ErrorResponse "Invalid input"
// @Failure 401 {object} model.ErrorResponse "Invalid email or password"
// @Router /login [post]
func (a *App) Login(c *gin.Context) {
	var newUser model.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "Bad request"})
		return
	}
	registeredUser, found := a.DB.GetUser(newUser.Email)
	if !found {
		c.IndentedJSON(http.StatusNotFound, model.ErrorResponse{Error: "Invalid email or password"})
		return
	}

	statusCode, err := a.MW.CompareHashAndPassword(registeredUser.Password, newUser.Password)
	if err != nil {
		c.IndentedJSON(statusCode, model.ErrorResponse{Error: err.Error()})
		return
	}

	tokenString, statusCode, err := a.MW.GenerateToken(registeredUser)
	if err != nil {
		c.IndentedJSON(statusCode, model.ErrorResponse{Error: err.Error()})
		return
	}
	registeredUser.Token = tokenString
	c.IndentedJSON(http.StatusOK, registeredUser)
}
