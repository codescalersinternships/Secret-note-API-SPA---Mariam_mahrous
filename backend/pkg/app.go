package app

import (
	"fmt"
	"net/http"
	"time"

	model "github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type App struct {
	R  *gin.Engine
	DB DB
}

func NewApp(db ConfigDB) *App {
	return &App{R: gin.Default(), DB: *NewDB(db)}
}

func (a *App) RegisterHandlers() {
	a.R.GET("/note/:uuid", a.GetNoteByUuid)
	a.R.GET("/note", a.RequireAuth, a.GetUserNotes)
	a.R.POST("/note/create", a.RequireAuth, a.CreateNote)
	a.R.GET("/validate", a.RequireAuth, a.validate)
	a.R.POST("/signup", a.SignUp)
	a.R.POST("/login", a.Login)
}

func (a *App) Run(port string) {
	a.RegisterHandlers()
	a.R.Run(":" + port)
}

//CODE MATCHINGGGGGG

func (a *App) GetNoteByUuid(c *gin.Context) {
	database := a.DB
	uuid, _ := uuid.Parse(c.Param("uuid"))
	note, statusCode := database.GetNoteByUuid(uuid)
	c.IndentedJSON(statusCode, note)
}

func (a *App) validate(c *gin.Context) {
	c.IndentedJSON(200, "logged in")
}

func (a *App) GetUserNotes(c *gin.Context) {
	id, ok := c.Get("id")
	if !ok {
		fmt.Printf("couldn't get user id")
		c.IndentedJSON(http.StatusUnauthorized, "Unauthorized access")
	}
	notes, statusCode := a.DB.GetAllNotes(id.(uint))
	c.IndentedJSON(statusCode, notes)
}

func (a *App) CreateNote(c *gin.Context) {
	var newNote model.Note
	if err := c.BindJSON(&newNote); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Check that all fields have been sent")
		return
	}
	database := a.DB
	id, ok := c.Get("id")
	if !ok {
		fmt.Printf("couldn't get user id")
		c.IndentedJSON(http.StatusUnauthorized, "Unauthorized access")
	}
	note, statusCode := database.CreateNote(id.(uint), newNote)
	c.IndentedJSON(statusCode, note)
}

func (a *App) SignUp(c *gin.Context) {
	var newUser model.User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	password := model.HashPassword(newUser.Password)
	newUser.Password = password
	statusCode := a.DB.SignUp(newUser)
	c.IndentedJSON(statusCode, newUser)
}

func (a *App) RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret-key"), nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		authorized, user := a.DB.VerifyUser(claims["email"])
		if authorized {
			c.Set("id", user.ID)
		}
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func (a *App) Login(c *gin.Context) {
	var newUser model.User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	registeredUser, found := a.DB.GetUser(newUser.Email)
	if !found {
		c.JSON(http.StatusNotFound, "Invalid email or password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(registeredUser.Password), []byte(newUser.Password))
	if err != nil {
		c.JSON(http.StatusNotFound, "Invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    registeredUser.ID,
		"email": registeredUser.Email,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	//Should i store this in .env and load it?
	tokenString, err := token.SignedString([]byte("secret-key"))

	if err != nil {
		c.IndentedJSON(http.StatusConflict, "Couldn't create Token")
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.IndentedJSON(http.StatusOK, "")
}
