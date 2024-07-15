package db

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email" validate:"email,required" gorm:"unique"`
	Token    string `json:"token"`
}

var validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func SignUp(c *gin.Context) {
	if c.Request.Method == "POST" {
		database := GetDatabaseSingelton().GetDatabase()
		var newUser User
		if err := c.BindJSON(&newUser); err != nil {
			return
		}

		validationErr := validate.Struct(newUser)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		password := HashPassword(newUser.Password)
		newUser.Password = password

		result := database.Where("email = ?", newUser.Email).First(&User{})
		if result.Error == nil {
			c.JSON(http.StatusInternalServerError, "E-Mail already exists")
		}

		if res := database.Create(&newUser); res.Error != nil {
			c.IndentedJSON(http.StatusConflict, "User can't be created")
			return
		}
		c.IndentedJSON(http.StatusOK, newUser)
	} else {
		c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func Login(c *gin.Context) {
	if c.Request.Method == "POST" {
		database := GetDatabaseSingelton().GetDatabase()
		var newUser User
		if err := c.BindJSON(&newUser); err != nil {
			return
		}

		validationErr := validate.Struct(newUser)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		var user User
		database.First(&user, "email = ?", newUser.Email)

		if user.ID == 0 {
			c.JSON(http.StatusNotFound, "Invalid email or password")

		}

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(newUser.Password))
		if err != nil {
			c.JSON(http.StatusNotFound, "Invalid email or password")
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": user.Email,
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
	} else {
		c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func Validate(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.IndentedJSON(http.StatusOK, "Logged in")
	} else {
		c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}
