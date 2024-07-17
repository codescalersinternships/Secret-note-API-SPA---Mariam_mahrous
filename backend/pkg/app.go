package app

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	model "github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type App struct {
	R  *gin.Engine
	DB DB
}

func NewApp(db ConfigDB) (*App, error) {
	database, err := NewDB(db)
	if err != nil {
		return &App{R: gin.Default(), DB: *database}, err
	}
	return &App{R: gin.Default(), DB: *database}, nil
}

func (a *App) RegisterHandlers() {
	a.R.GET("/note/:uuid", a.GetNoteByUuid)
	a.R.GET("/note", a.RequireAuth, a.GetUserNotes)
	a.R.POST("/note/create", a.RequireAuth, a.CreateNote)
	a.R.POST("/signup", a.SignUp)
	a.R.POST("/login", a.Login)
}

func (a *App) Run(port string) error {

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
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
	if !ok {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}
	notes, statusCode, err := a.DB.GetAllNotes(id.(uint))
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
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

	password, err := model.HashPassword(newUser.Password)
	if err != nil {
		c.IndentedJSON(http.StatusConflict, gin.H{"error": "Error hashing the password"})
		return
	}
	newUser.Password = password
	tokenString , statusCode , err := a.generateToken(newUser)
	newUser.Token = tokenString
	if err !=nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
	}
	statusCode, err = a.DB.SignUp(newUser)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(statusCode, newUser)
}

func (a *App) RequireAuth(c *gin.Context) {
	tokenString :=  strings.Split(c.Request.Header["Authorization"][0], " ")[1]
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
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Invalid email or password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(registeredUser.Password), []byte(newUser.Password))
	if err != nil {
		c.JSON(http.StatusNotFound, "Invalid email or password")
		return
	}
	tokenString , statusCode , err := a.generateToken(registeredUser)
	if err !=nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
	}
	registeredUser.Token=tokenString
	//save user to db
	c.IndentedJSON(http.StatusOK, registeredUser)
}

func (a*App) generateToken(user model.User) (string,int ,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	//Should i store this in .env and load it?
	tokenString, err := token.SignedString([]byte("secret-key"))

	if err != nil {
		return "", http.StatusConflict , errors.New("couldn't create Token")
	}
	return tokenString , 0 , nil
}